#!/bin/bash
# Deploy provider-libvirt to cluster

set -e

# Configuration - update these for your environment
REGISTRY="${REGISTRY:-ghcr.io/nourspeed}"  # or your private registry
IMAGE_NAME="provider-libvirt"
VERSION="v0.0.0-5.ga53b606.dirty"
NAMESPACE="${NAMESPACE:-crossplane-system}"

echo "=== Step 1: Push Docker image to registry ==="
# Tag the local image
docker tag build-d0e382f3/provider-libvirt-amd64:latest ${REGISTRY}/${IMAGE_NAME}:${VERSION}
docker tag build-d0e382f3/provider-libvirt-amd64:latest ${REGISTRY}/${IMAGE_NAME}:latest

# Push to registry (ensure you're logged in first)
echo "Pushing to ${REGISTRY}/${IMAGE_NAME}:${VERSION}"
docker push ${REGISTRY}/${IMAGE_NAME}:${VERSION}
docker push ${REGISTRY}/${IMAGE_NAME}:latest

echo "=== Step 2: Check Crossplane installation ==="
if ! kubectl get deployment crossplane -n ${NAMESPACE} &>/dev/null; then
    echo "Crossplane not found. Installing Crossplane..."
    kubectl create namespace ${NAMESPACE} || true
    helm repo add crossplane-stable https://charts.crossplane.io/stable
    helm repo update
    helm install crossplane \
        --namespace ${NAMESPACE} \
        crossplane-stable/crossplane \
        --set args='{"--enable-external-secret-stores"}'
    
    echo "Waiting for Crossplane to be ready..."
    kubectl wait --for=condition=available --timeout=300s deployment/crossplane -n ${NAMESPACE}
else
    echo "Crossplane is already installed"
fi

echo "=== Step 3: Create Provider manifest ==="
cat > provider-libvirt.yaml <<EOF
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-libvirt
spec:
  package: ${REGISTRY}/${IMAGE_NAME}:${VERSION}
  controllerConfigRef:
    name: provider-libvirt-config
---
apiVersion: pkg.crossplane.io/v1alpha1
kind: ControllerConfig
metadata:
  name: provider-libvirt-config
spec:
  args:
    - --debug
  env:
    - name: TERRAFORM_VERSION
      value: "1.2.1"
EOF

echo "=== Step 4: Install the Provider ==="
kubectl apply -f provider-libvirt.yaml

echo "Waiting for provider to be healthy..."
kubectl wait --for=condition=healthy --timeout=300s provider.pkg.crossplane.io/provider-libvirt

echo "=== Step 5: Create ProviderConfig ==="
# Update this with your libvirt connection details
cat > providerconfig.yaml <<EOF
apiVersion: libvirt.nourspeed.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: ${NAMESPACE}
      name: libvirt-credentials
      key: uri
EOF

# Create secret with libvirt URI
read -p "Enter your libvirt URI (e.g., qemu+ssh://user@host/system): " LIBVIRT_URI
kubectl create secret generic libvirt-credentials \
  -n ${NAMESPACE} \
  --from-literal=uri="${LIBVIRT_URI}" \
  --dry-run=client -o yaml | kubectl apply -f -

kubectl apply -f providerconfig.yaml

echo "=== Deployment Complete ==="
echo "Provider status:"
kubectl get providers
kubectl get pods -n ${NAMESPACE} | grep provider-libvirt

echo ""
echo "To test the provider, create a resource like:"
echo "kubectl apply -f examples/network/network.yaml"