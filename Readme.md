# Fiber Name
Fiber Name is a api service to load your name

### Requirement API
- Golang 1.21.1
- Go Fiber V2
- Docker

### CICD Github Actions

##### Secret
- DOCKER_USERNAME => Docker username
- DOCKER_PASSWORD => Docker Password
- USERNAME => VM username
- PRIVATE_KEY => VM private key

##### Variable
- HOST => VM host
- PATH  => PATH to place dockerfile
- IMAGE => Image for container

### How to  Install
##### Docker
Run service using docker 
```sh
docker compose up
```

** set image repo on docker-compose.yaml
##### Kubernetes
Run service using docker 
```sh
kubectl apply -f ./manifest/
```
