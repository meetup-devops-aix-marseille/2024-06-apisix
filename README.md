# meetup apisix

## werf to deploy

If you are using buildah with werf

```bash
export WERF_BUILDAH_MODE=auto
```

```bash
werf converge --repo ttl.sh/meetup-apisix --dev
```

## Setup with HELM

### Links

[helm chart](https://charts.apiseven.com)
[apisix](https://artifacthub.io/packages/helm/apisix/apisix)
[apisix-ingress-controller](https://artifacthub.io/packages/helm/apisix/apisix-ingress-controller)
[apisix-dashboard](https://artifacthub.io/packages/helm/apisix/apisix-dashboard)

### APISIX

```bash
helm upgrade --install --repo https://charts.apiseven.com apisix apisix --namespace apisix --create-namespace  -f helm/values.yaml
```
