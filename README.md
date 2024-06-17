# meetup apisix

## Setup with HELM

### APISIX

```bash
helm upgrade --install --repo https://charts.apiseven.com apisix apisix --namespace apisix --create-namespace  -f helm/values.yaml
```
