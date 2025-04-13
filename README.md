[![Helm Checks](https://github.com/Criyl/portfolio-operator/actions/workflows/ci.yaml/badge.svg)](https://github.com/Criyl/portfolio-operator/actions/workflows/ci.yaml)
# Portfolio Kubernetes Operator
Manage your portfolio natively in your kubernetes cluster.

### Install with Helm
Chart is available as an image on [ghcr](https://github.com/Criyl/portfolio-operator/pkgs/container/portfolio-operator)
```bash
helm install portfolio-operator oci://ghcr.io/criyl/portfolio-operator
```

### Supported Annotations
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
```

### Environment Variables
| KEY        | DESCRIPTION               | DEFAULT        |
| ---------- | ------------------------- | -------------- |
| KUBECONFIG | Path to kubernetes config | ~/.kube/config |
| DEBUG      | Deploy with swagger UI    | true           |
| HOST       | Host to serve CRUD api    | 0.0.0.0        |
| PORT       | Port to serve CRUD api    | 8080           |