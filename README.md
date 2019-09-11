# EMQ X example deployment to K8S

The goal of this repo is to configure an EMQ cluster running on kubernetes.

## Janky dev setup

Database service is tied to minikube currently. And probably have to set `MQTT_URL` env to your cluster ip.

-   `docker-compose up db -d`
-   `kubectl apply -f database-service.yml`
-   `kubectl apply -f emq-service.yml`

Also you can run the app in cluster

-   `kubectl apply -f app-deployment`

## TLS Setup

Self signed TLS is enabled by default, set enableTLS to false in client options to use unsecured connection.

### Key setup

-   Gen self-signed CA root cert

    `cd app/certs && openssl genrsa -out emqxca.key 2048`

-   Gen root cert

    `openssl req -x509 -new -nodes -key emqxca.key -sha256 -days 3650 -out emqxca.pem`

    _The root cert is the start of a trust chain, we assume that a cert on the chain is trustworthy if all the nodes from its issuer up to the root cert on the chain are trustworthy._

-   After having a self-signed Root CA Certificate, we can use it to issue certificates for other identities, like the EMQ server. Similarly, we will need a private key first

    `openssl genrsa -out emqx.key 2048`

-   Then the certificate request for emqx

    `openssl req -new -key ./emqx.key -out emqx.csr`

-   Then use the root CA to issue the cert for emqx

    `openssl x509 -req -in ./emqx.csr -CA emqxca.pem -CAkey emqxca.key -CAcreateserial -out emqx.pem -days 3650 -sha256`

-   Finally update the secret in the emq service yaml

```
 ---
apiVersion: v1
kind: Secret
metadata:
 name: cert-secrets
data:
 key.pem: #Here
 cert.pem: #Here
 cacert.pem: #Here
---
```
