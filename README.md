# About

Start a minikbe cluster:
```
minikube start \
  --memory=4096 \
  --cpus=2 \
  --disk-size=20g \
  --kubernetes-version=v1.30.0 \
  --driver=docker \
  --profile=scaling-ws
```


Build services using Tilt:
```
tilt up
```