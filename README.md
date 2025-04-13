[![Helm Checks](https://github.com/Criyl/portfolio-operator/actions/workflows/ci.yaml/badge.svg)](https://github.com/Criyl/portfolio-operator/actions/workflows/ci.yaml)

# Portfolio Kubernetes Operator

Manage your portfolio natively in your kubernetes cluster.

### What is it?
With the Portfolio Operator you can define the state of your portfolio dynamically using the Portfolio Custom Resource. Portfolio's will also be automatically created by annotating ingresses with the [supported annotations](#supported-ingress-annotations). 

The spec of these Portfolio's is conveniently available through a CRUD api so it can be queried for your portfolio website without exposing your cluster. 

### Install with Helm
The chart is available as an image on [ghcr](https://github.com/Criyl/portfolio-operator/pkgs/container/portfolio-operator)
```bash
helm install portfolio-operator oci://ghcr.io/criyl/portfolio-operator
```

### Supported Ingress Annotations
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: ingress
  annotations:
    portfolio-operator/name: "portfolio-operator"
    portfolio-operator/url: "https://portfolio-operator.carroll.codes"
    portfolio-operator/blog: "https://github.com/Criyl/portfolio-operator"
    portfolio-operator/icon: "https://portfolio-operator.carroll.codes/swagger/favicon-16x16.png"
    portfolio-operator/healthcheck: "https://portfolio-operator.carroll.codes/health"
    portfolio-operator/tags: "tag1,tag2,tag3"
...
```

### Portfolio CRD
```yaml
apiVersion: carroll.codes/v1
kind: Portfolio
metadata:
  name: portfolio-portfolio-ingress
spec:
  name: portfolio-operator
  url: https://portfolio-operator.carroll.codes
  blog: https://github.com/Criyl/portfolio-operator
  icon: https://portfolio-operator.carroll.codes/swagger/favicon-16x16.png
  healthcheck: https://portfolio-operator.carroll.codes/health
  tags:
  - tag1
  - tag2
  - tag3
```

## Environment Variables

| KEY          | DESCRIPTION                          | DEFAULT        |
| ------------ | ------------------------------------ | -------------- |
| KUBECONFIG   | Path to kubernetes config            | ~/.kube/config |
| DEBUG        | Deploy with swagger UI               | true           |
| HOST         | Host to serve CRUD api               | 0.0.0.0        |
| PORT         | alias for API_PORT                   |                |
| API_PORT     | Port to serve CRUD api               | 8080           |
| METRICS_PORT | Port to serve Controller Metrics api | 8081           |
