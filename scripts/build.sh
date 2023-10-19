#!/bin/bash

kind create cluster --config deployments/kind.yml
kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:acvii
docker build -t achoistic98/state:v0.0.1 -f build/Dockerfile.acvii .
docker push achoistic98/acvii:v0.0.1
kubectl apply -f deployments/acvii.yml