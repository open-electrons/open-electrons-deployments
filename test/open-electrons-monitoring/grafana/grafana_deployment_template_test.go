package grafana

import (
	"testing"

	appsv1 "k8s.io/api/apps/v1"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/assert"
	//"github.com/gruntwork-io/terratest/modules/k8s"
	//"github.com/gruntwork-io/terratest/modules/logger"
)

func TestGrafanaHelmChartTemplate(t *testing.T) {
	// Path to the helm chart we will test
	helmChartGrafanaPath := "../../../open-electrons-monitoring"

	// Setup the args. For this test, we will set the following input values:
	// - image=grafana:latest
	options := &helm.Options{
		SetValues: map[string]string{
			"image": "grafana/grafana:latest",
			"name": "open-electrons-grafana",
			"namespace": "open-electrons-monitoring-ns",
		},
	}

	// Run RenderTemplate to render the template and capture the output.
	output := helm.RenderTemplate(t, options, helmChartGrafanaPath, "deployment", []string{"templates/grafana/grafana-deployment.yml"})

	// Now we use kubernetes/client-go library to render the template output into the Pod struct. This will
	// ensure the Pod resource is rendered correctly.
	var deployment appsv1.Deployment
	helm.UnmarshalK8SYaml(t, output, &deployment)

	// Finally, we verify the pod spec is set to the expected container image value
	expectedContainerImage := "grafana/grafana:latest"
	expectedNamespace := "open-electrons-monitoring-ns"
	podContainers := deployment.Spec.Template.Spec.Containers
	assert.Equal(t, expectedContainerImage, podContainers[0].Image, "Expected container image did not match")
	assert.Equal(t, expectedNamespace, deployment.Namespace)
}