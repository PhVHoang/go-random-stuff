#!/usr/bin/env

kubectl run demo --image=cloudnatived/demo: hello --port=9999 --labels app=demo
kubectl port-forward deploy/demo 9999:8888

# Check the pod is indeed running
kubectl get pods --selector app=demo

# Stop the pod
kubectl delete pods --selector app=demo

# And list the pods again
kubectl get pods --selector app=demo

# shut it down and clean up
kubectl delete all --selector app=demo

# run the yaml resource file
kubectl apply -f deployment.yaml

# create the Service
kubectl apply -f service.yaml

kubectl port-forward service/demo 9999:8888

# See comprehensive information about an individual Pod
kubectl describe pod/<pod_name>

kubectl delete -f ./

