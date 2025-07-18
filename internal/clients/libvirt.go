/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/nourspeed/provider-libvirt/apis/v1beta1"
)

const (
	keyUri      = "uri"
	keyCACert   = "cacert"     // CA certificate content
	keyPKIPath  = "pkipath"    // Custom PKI directory path
	keyNoVerify = "no_verify"  // Skip certificate verification
	
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal libvirt credentials as JSON"
	errSetupCustomCA        = "cannot setup custom CA certificate"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials in Terraform provider configuration.
		ps.Configuration = map[string]any{}
		if v, ok := creds[keyUri]; ok {
			ps.Configuration[keyUri] = v
		}

		// Handle custom CA certificate
		if caCert, ok := creds[keyCACert]; ok && caCert != "" {
			pkiPath := "/tmp/libvirt-pki"
			if err := setupCustomCA(pkiPath, caCert); err != nil {
				return ps, errors.Wrap(err, errSetupCustomCA)
			}
			ps.Configuration[keyPKIPath] = pkiPath
		}

		// Handle custom PKI path (if specified and not using inline CA)
		if pkiPath, ok := creds[keyPKIPath]; ok && pkiPath != "" {
			// Only use if no inline CA was provided
			if _, hasCACert := creds[keyCACert]; !hasCACert {
				ps.Configuration[keyPKIPath] = pkiPath
			}
		}

		// Handle certificate verification options
		if noVerify, ok := creds[keyNoVerify]; ok {
			ps.Configuration[keyNoVerify] = noVerify
		}

		return ps, nil
	}
}

// setupCustomCA creates a PKI directory with the provided CA certificate
func setupCustomCA(pkiPath, caCert string) error {
	// Create PKI directory
	if err := os.MkdirAll(pkiPath, 0755); err != nil {
		return errors.Wrap(err, "cannot create PKI directory")
	}

	// Write CA certificate
	caPath := filepath.Join(pkiPath, "cacert.pem")
	if err := ioutil.WriteFile(caPath, []byte(caCert), 0644); err != nil {
		return errors.Wrap(err, "cannot write CA certificate")
	}

	return nil
}