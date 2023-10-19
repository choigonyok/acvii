# ACViI

ACViI is abbreviation of "Access Control Visualizer in Istio". ACViI helps verifying current state of Istio AuthorizationPolicy, maintaining desired state as current state of Istio AuthorizationPolicy.

## Dependencies

Kubectl (v1.28.1)
Istio (v1.19)

## Prerequisites

To run ACViI, you need Istio deployed in running Kubernetes cluster.

## Quickstart

Run the following command to deploy ACViI in your Kubernetes cluster:

```
kubectl apply -f https://github.com/choigonyok/acvii/releases/download/latest/acvii.yml
```

## Documentation

## Contributing