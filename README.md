## a-micro
````
docker build -t a-micro .
docker run -p 8080:8080 a-micro
````

### Docker push
````
docker build -t mytrgor/a-micro ./a-micro
docker push mytrgor/a-micro
````

## b-micro
````
docker build -t b-micro .
docker run -p 8081:8081 b-micro
````

### Docker push
```
docker build -t b-micro .
docker tag b-micro:latest mytrgor/b-micro:v1.0.3
docker push mytrgor/b-micro:v1.0.3
```

## docker-compose
````
docker-compose up --build
````

## eksctl
````
eksctl create cluster --name k8s-micros-002 --region us-east-1 --zones us-east-1a,us-east-1b
````

````
eksctl delete cluster --name k8s-micros-001
````

## kubectl


kubectl apply -f b.yaml


````
kubectl apply -f b-micro-deployment.yaml
kubectl apply -f b-micro-service.yaml
kubectl apply -f b-micro-hpa.yaml
````

```
kubectl delete hpa b-micro-hpa
kubectl delete svc b-micro
kubectl delete deployment b-micro
```


````
kubectl apply -f a-micro-deployment.yaml
kubectl apply -f a-micro-service.yaml
kubectl apply -f a-micro-hpa.yaml
````



kubectl exec -it b-micro-deployment-76c4b689d6-csbm7 -c e788051ebd4f42b93edfd159f6aef11b0360e3655939c7fe65166ebdf418fd94 -- /bin/bash


kubectl describe pod b-micro-deployment-76c4b689d6-csbm7

docker pull mytrgor/b-micro
docker run -d -p 8081:8081 mytrgor/b-micro

kubectl logs b-micro-6b776989c6-lnww6

kubectl exec -it b-micro-deployment-76c4b689d6-csbm7 -- /bin/sh

