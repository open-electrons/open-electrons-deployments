## test

TODO: Under documentation

Just like how we test code, we can also test code here (maybe structured code, I mean the kubernetes yaml files). Just like
how we test our code in a CI pipeline, we do the same here with GitHub Actions. Have a look [here](https://github.com/open-electrons/open-electrons-deployments/.github/workflows) where the CI pipeline kicks in
for every commit that happens in a feature branch.

We use GitHub Actions for our CI, and use the following tools:

1. Kube Tools - https://github.com/stefanprodan/kube-tools

2. Conftest - https://github.com/instrumenta/conftest

3. kubeval - https://github.com/instrumenta/kubeval

You could ofcourse run the conftest and kubeval locally on your machine before pushing your changes into the repository.

### Running Conftest

TODO: Document usage

To run conftest locally via helm, an additional helm plugin need to be installed (Assuming that you have already installed helm locally)

- Install the helm-conftest plugin from here - https://github.com/instrumenta/helm-conftest

Conftest uses Open Policy Agent for writing the tests in the Rego language. A nice place to play with it is the playground here:

- https://play.openpolicyagent.org/p/qwwAdgbc6S

### Running kubeval

Navigate to the root folder of this project and issue the following command:

Using Helm:

```
helm template helm-k8s | kubeval --strict --force-color --ignore-missing-schemas
```

or simply

```
kubeval -d helm-k8s/templates/base,helm-k8s/templates/dev,helm-k8s/templates/production --force-color --strict --ignore-missing-schemas
```

