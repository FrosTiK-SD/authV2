#!/bin/sh

# Connect to minikube docker env
eval $(minikube docker-env)
echo Connected to minikube docker env

# Build the image
echo Building the image
docker build -t authv2 -f Dockerfile . --build-arg ATLAS_URI=$ATLAS_URI --build-arg FIREBASE_PROJECT_ID=$FIREBASE_PROJECT_ID
echo Docker image built successfully

# Deploy redis
kubectl create deployment redis --image=redis/redis-stack-server:latest
kubectl expose deployment redis --port=6379   

# Deploy the image
echo Deleteing the older deployments
kubectl delete deployment authv2 --force --grace-period=0
echo Deploying the image
kubectl apply -f ./deploy/deployment.yaml
echo Deployment Done !!!

# Run it
echo Now run the server by
echo        make serve

