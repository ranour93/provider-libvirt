package clients

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/crossplane-runtime/pkg/test"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/nourspeed/provider-libvirt/apis/v1beta1"
)

func TestTerraformSetupBuilder_CustomCA(t *testing.T) {
	type args struct {
		version         string
		providerSource  string
		providerVersion string
	}
	type want struct {
		err    error
		config map[string]any
	}

	testCA := "-----BEGIN CERTIFICATE-----\nMIIC...TEST CA CERTIFICATE...==\n-----END CERTIFICATE-----"

	cases := map[string]struct {
		args args
		mg   resource.Managed
		pc   *v1beta1.ProviderConfig
		want want
	}{
		"BasicURI": {
			args: args{
				version:         "1.6.6",
				providerSource:  "dmacvicar/libvirt",
				providerVersion: "0.7.6",
			},
			mg: &mockManaged{
				ref: &xpv1.Reference{Name: "test-config"},
			},
			pc: &v1beta1.ProviderConfig{
				Spec: v1beta1.ProviderConfigSpec{
					Credentials: v1beta1.ProviderCredentials{
						Source: xpv1.CredentialsSourceSecret,
						CommonCredentialSelectors: xpv1.CommonCredentialSelectors{
							SecretRef: &xpv1.SecretKeySelector{
								SecretReference: xpv1.SecretReference{
									Name:      "test-secret",
									Namespace: "test-namespace",
								},
								Key: "credentials",
							},
						},
					},
				},
			},
			want: want{
				config: map[string]any{
					"uri": "qemu+tls://test.example.com/system",
				},
			},
		},
		"CustomCA": {
			args: args{
				version:         "1.6.6",
				providerSource:  "dmacvicar/libvirt",
				providerVersion: "0.7.6",
			},
			mg: &mockManaged{
				ref: &xpv1.Reference{Name: "test-config"},
			},
			pc: &v1beta1.ProviderConfig{
				Spec: v1beta1.ProviderConfigSpec{
					Credentials: v1beta1.ProviderCredentials{
						Source: xpv1.CredentialsSourceSecret,
						CommonCredentialSelectors: xpv1.CommonCredentialSelectors{
							SecretRef: &xpv1.SecretKeySelector{
								SecretReference: xpv1.SecretReference{
									Name:      "test-secret",
									Namespace: "test-namespace",
								},
								Key: "credentials",
							},
						},
					},
				},
			},
			want: want{
				config: map[string]any{
					"uri":     "qemu+tls://test.example.com/system",
					"pkipath": "/tmp/libvirt-pki",
				},
			},
		},
		"NoVerifyOption": {
			args: args{
				version:         "1.6.6",
				providerSource:  "dmacvicar/libvirt",
				providerVersion: "0.7.6",
			},
			mg: &mockManaged{
				ref: &xpv1.Reference{Name: "test-config"},
			},
			pc: &v1beta1.ProviderConfig{
				Spec: v1beta1.ProviderConfigSpec{
					Credentials: v1beta1.ProviderCredentials{
						Source: xpv1.CredentialsSourceSecret,
						CommonCredentialSelectors: xpv1.CommonCredentialSelectors{
							SecretRef: &xpv1.SecretKeySelector{
								SecretReference: xpv1.SecretReference{
									Name:      "test-secret",
									Namespace: "test-namespace",
								},
								Key: "credentials",
							},
						},
					},
				},
			},
			want: want{
				config: map[string]any{
					"uri":       "qemu+tls://test.example.com/system",
					"no_verify": "true",
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			// Setup mock credentials based on test case
			var credData map[string]string
			switch name {
			case "BasicURI":
				credData = map[string]string{
					"uri": "qemu+tls://test.example.com/system",
				}
			case "CustomCA":
				credData = map[string]string{
					"uri":    "qemu+tls://test.example.com/system",
					"cacert": testCA,
				}
			case "NoVerifyOption":
				credData = map[string]string{
					"uri":       "qemu+tls://test.example.com/system",
					"no_verify": "true",
				}
			}

			credJSON, _ := json.Marshal(credData)

			// Create mock client
			client := test.NewMockClient()
			client.MockGet = test.NewMockGetFn(nil, func(obj client.Object) error {
				switch o := obj.(type) {
				case *v1beta1.ProviderConfig:
					*o = *tc.pc
					return nil
				case *corev1.Secret:
					o.Data = map[string][]byte{
						"credentials": credJSON,
					}
					return nil
				}
				return nil
			})

			// Create setup function and execute
			setupFn := TerraformSetupBuilder(tc.args.version, tc.args.providerSource, tc.args.providerVersion)
			got, err := setupFn(context.Background(), client, tc.mg)

			// Verify results
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("TerraformSetupBuilder() error = %v, want %v\n%s", err, tc.want.err, diff)
				return
			}

			if err != nil {
				return
			}

			if diff := cmp.Diff(tc.want.config, got.Configuration); diff != "" {
				t.Errorf("TerraformSetupBuilder() config = %v, want %v\n%s", got.Configuration, tc.want.config, diff)
			}

			// For CustomCA test, verify the CA file was created
			if name == "CustomCA" {
				pkiPath := "/tmp/libvirt-pki"
				caPath := filepath.Join(pkiPath, "cacert.pem")
				if _, err := os.Stat(caPath); os.IsNotExist(err) {
					t.Errorf("CA certificate file was not created at %s", caPath)
				} else {
					// Verify content
					content, err := ioutil.ReadFile(caPath)
					if err != nil {
						t.Errorf("Cannot read CA certificate: %v", err)
					} else if string(content) != testCA {
						t.Errorf("CA certificate content mismatch")
					}
				}
				// Cleanup
				os.RemoveAll(pkiPath)
			}
		})
	}
}

func TestSetupCustomCA(t *testing.T) {
	testCA := "-----BEGIN CERTIFICATE-----\nTEST CA CONTENT\n-----END CERTIFICATE-----"

	// Create temporary directory for testing
	tmpDir, err := ioutil.TempDir("", "libvirt-ca-test-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Test setupCustomCA function
	err = setupCustomCA(tmpDir, testCA)
	if err != nil {
		t.Errorf("setupCustomCA() error = %v", err)
		return
	}

	// Verify CA file was created with correct content
	caPath := filepath.Join(tmpDir, "cacert.pem")
	content, err := ioutil.ReadFile(caPath)
	if err != nil {
		t.Errorf("Cannot read CA file: %v", err)
		return
	}

	if string(content) != testCA {
		t.Errorf("CA file content = %q, want %q", string(content), testCA)
	}

	// Verify file permissions
	info, err := os.Stat(caPath)
	if err != nil {
		t.Errorf("Cannot stat CA file: %v", err)
		return
	}

	expectedMode := os.FileMode(0644)
	if info.Mode().Perm() != expectedMode {
		t.Errorf("CA file mode = %v, want %v", info.Mode().Perm(), expectedMode)
	}
}

// Mock implementation for testing
type mockManaged struct {
	ref *xpv1.Reference
}

func (m *mockManaged) GetProviderConfigReference() *xpv1.Reference { return m.ref }
func (m *mockManaged) SetProviderConfigReference(r *xpv1.Reference) { m.ref = r }
func (m *mockManaged) GetDeletionPolicy() xpv1.DeletionPolicy { return xpv1.DeletionDelete }
func (m *mockManaged) SetDeletionPolicy(p xpv1.DeletionPolicy) {}
func (m *mockManaged) GetManagementPolicies() xpv1.ManagementPolicies { return nil }
func (m *mockManaged) SetManagementPolicies(p xpv1.ManagementPolicies) {}
func (m *mockManaged) GetProviderReference() *xpv1.Reference { return nil }
func (m *mockManaged) SetProviderReference(r *xpv1.Reference) {}
func (m *mockManaged) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo { return nil }
func (m *mockManaged) SetPublishConnectionDetailsTo(p *xpv1.PublishConnectionDetailsTo) {}
func (m *mockManaged) GetWriteConnectionSecretsToReference() *xpv1.SecretReference { return nil }
func (m *mockManaged) SetWriteConnectionSecretsToReference(r *xpv1.SecretReference) {}
func (m *mockManaged) GetCondition(ct xpv1.ConditionType) xpv1.Condition { return xpv1.Condition{} }
func (m *mockManaged) SetConditions(c ...xpv1.Condition) {}
func (m *mockManaged) GetObjectKind() runtime.ObjectKind { return nil }
func (m *mockManaged) DeepCopyObject() runtime.Object { return m }
func (m *mockManaged) GetObjectMeta() metav1.Object { return &metav1.ObjectMeta{} }