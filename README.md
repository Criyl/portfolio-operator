# Portfolio Kubernetes Operator
Manage your portfolio natively in your kubernetes cluster.

## Features
- Query portfolio metadata.
- Dynamically update your portfolio by annotating ingresses.
- Link blog posts to your ingresses.

### Supported Annotations
```yaml
metadata:
  annotations:
  - portfolio-operator/name: ""
  - portfolio-operator/url: ""
  - portfolio-operator/blog: ""
  - portfolio-operator/icon: ""
  - portfolio-operator/healthcheck: ""
  - portfolio-operator/tags: "tag1,tag2,tag3"
```

### Planned Features
- Pagination
- Adding Dates to the schema