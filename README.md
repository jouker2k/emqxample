# EMQ X example deployment to K8S

The goal of this repo is to configure an EMQ cluster running on kubernetes.

## Janky dev setup

Database service is tied to minikube currently. And probably have to set `MQTT_URL` env to your cluster ip.

-   `docker-compose up db -d`
-   `kubectl apply -f database-service.yml`
-   `kubectl apply -f emq-service.yml`

Also you can run the app in cluster

-   `kubectl apply -f app-deployment`

## TODO

-   Get logging to provide better overview of comms
