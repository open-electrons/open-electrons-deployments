# open-electrons-deployments

TODO: Currently under construction

Contains the Kubernetes deployment manifests for all open-electrons projects

## helm

install helm unittests using:

´´´
helm plugin install https://github.com/quintush/helm-unittest
´´´

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

## unit-test

Like everything, we unit test and continuously integrate the unit tests with a GitHub Actions pipeline. To run
the unit tests for the helm charts (we only do template testing as of now), navigate to any project folder and run

´´´
go test -v
´´´