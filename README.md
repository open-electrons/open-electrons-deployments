# open-electrons-deployments

TODO: Under construction

Contains the Kubernetes deployment manifests for all open-electrons projects

## helm
We are using helm for templating & packaging. We use the standard helm folder structure:

```
├── charts
├── templates
├   ├── base
│   ├    ├── application
│   ├        ├── xxx-deployment.yaml
│   ├        ├── xxx-namespace.yaml
│   ├        └── xxx-service.yaml
│   ├── kustomization.yaml
│   ├    ├── monitoring
│   ├        ├── todo.yaml
│   ├        ├── todo.yaml
│   ├        ├── todo.yaml
│   ├        └── todo.yaml
├   ├── test
│   ├    ├── flux-patch.yaml
│   ├    └── kustomization.yaml
│   ├── prod
├   ├    ├── flux-patch.yaml
└   ├    ├── kustomization.yaml
```
