# open-electrons-deployments

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
├   ├── dev
│   ├    ├── flux-patch.yaml
│   ├    └── kustomization.yaml
│   ├── production
├   ├    ├── flux-patch.yaml
└   ├    ├── kustomization.yaml
```
