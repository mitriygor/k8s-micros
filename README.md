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
docker build -t mytrgor/b-micro ./b-micro
docker push mytrgor/b-micro
```

## docker-compose
````
docker-compose up --build
````

## eksctl
````
eksctl create cluster --name k8s-micros-001 --region us-east-1 --zones us-east-1a,us-east-1b --node-type t3.micro --nodes 2
````

## kubectl

````
kubectl apply -f b-micro-service.yaml
kubectl apply -f b-micro-deployment.yaml
kubectl apply -f b-micro-hpa.yaml
````

````
kubectl apply -f a-micro-deployment.yaml
kubectl apply -f a-micro-service.yaml
kubectl apply -f a-micro-hpa.yaml
````
