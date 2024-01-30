gcloud container clusters get-credentials cluster-1 --zone=us-central1-c


docker build -t myapp .

docker tag myapp gcr.io/playground-s-11-c6766f2d/myapp:blue
docker push gcr.io/playground-s-11-c6766f2d/myapp:blue

docker tag myapp us-central1-docker.pkg.dev/playground-s-11-c6766f2d/myapp-repository/myapp:blue
docker push us-central1-docker.pkg.dev/playground-s-11-c6766f2d/myapp-repository/myapp:blue

#for kubectl
gcloud container clusters get-credentials cluster-1 --zone us-central1-c --project playground-s-11-c6766f2d

kubectl apply -f myapp-deployment.yaml 
kubectl apply -f myapp-service.yaml 
kubectl get pods
kubectl get deployments
kubectl get services
kubectl get events

kubectl describe deployment myapp-deployment

kubectl expose deployment nginx --port=80 --type=LoadBalancer

kubectl logs pods/name


