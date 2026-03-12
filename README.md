# hades-artemis-adapter

> [!WARNING]
> This is a work in progress repository.

## Deployment

```shell

helm upgrade --install hades-artemis-adapter ./helm/hades-artemis-adapter -f ./helm/values-deploy.yaml \
    --namespace hades-artemis-connector --create-namespace \
    --set image.tag="main" \
    --set env.ARTEMIS_BASE_URL="https://artemis-test8.artemis.cit.tum.de" \
    --set env.ARTEMIS_AUTH_TOKEN="<TOKEN>"

```
