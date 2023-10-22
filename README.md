# ACViI

**ACViI** is abbreviation of **"Access Control Visualizer in Istio"**</br>
ACViI helps verifying current state of Istio AuthorizationPolicy with dashboard UI, maintaining desired state as current state of Istio AuthorizationPolicy.

## Dependencies

kubectl (v1.28.1)
Istio (v1.19)

## Prerequisites

To run ACViI, you need Istio deployed in running Kubernetes cluster.

## Quickstart with acvctl

**acvctl** is **CLI** (Command Line Interface) for ACViI.

Run the following command to deploy ACViI pod in your Kubernetes cluster:
```
acvctl install -y
```

You can check what's gonna be different before you apply AuthorizationPolicy with:
```
acvctl plan
``` 

You can apply your AuthorizationPolicies with:
```
acvctl apply
``` 

Run the following command to verify current/desired state of AuthorizationPolicies on UI dashboard:
```
acvctl dashboard
``` 

You can kill running acvii pod with:
```
acvctl uninstall -y
``` 

Running the following command will help you to use acvctl:
```
acvctl --help
``` 

## Documentation

You can check docs about ACViI in 
[HERE](https://github.com/choigonyok/acvii/tree/main/docs)

## Contributing

You can check about contributing in 
[HERE](https://github.com/choigonyok/acvii/blob/main/CONTRIBUTING.md)