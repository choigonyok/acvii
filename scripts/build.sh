#!/bin/bash

kind create cluster --config deployments/kind.yml
kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:goopt
docker build -t achoistic98/goopt:v0.0.1 -f build/Dockerfile .
docker push achoistic98/goopt:v0.0.1
kubectl apply -f deployments/goopt.yml