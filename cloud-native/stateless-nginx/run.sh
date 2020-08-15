#!/usr/bin/env

kubectl apply -f k8s/deployment.yaml

kubectl describe deployment nginx-deployment

kubectl get pods -l app=nginx

# delete this deployment

kubectl delete deployment nginx-deployment
