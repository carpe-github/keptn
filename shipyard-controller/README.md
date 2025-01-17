# Shipyard Controller

## Installation

The *shipyard-controller* is installed as a part of [keptn](https://keptn.sh)

## Deploy in your Kubernetes cluster

This service should be automatically deployed when executing `keptn install` or installing Keptn from a Helm chart. If
you still want to deploy it manually in your Keptn Kubernetes, you can either

* use the manifest `deploy/service.yaml` from this repository and apply it
  ```console
  kubectl apply -f deploy/service.yaml
  ```
* build locally (from source) and deploy it using
  ```console
  export SKAFFOLD_DEFAULT_REPO=containerregistry # e.g., docker.io/username
  skaffold run --tail  
  ```
  **Note**: When stopping skaffold (e.g., using CTRL-C) the deployment will automatically be cleaned up

## Delete in your Kubernetes cluster

To delete a deployed *shipyard-controller*, use the manifest `deploy/service.yaml` from this repository and delete the
Kubernetes resources:

```console
kubectl delete -f deploy/service.yaml
```

## Generate  Swagger doc from source

1. Download and install Swag for Go by calling `go get -u github.com/swaggo/swag/cmd/swag` in fresh terminal.
2. `cd` to the Shipyard Controller's root folder and run `swag init --parseDependency`
