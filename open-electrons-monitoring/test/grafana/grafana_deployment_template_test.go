package grafana

import (
	"testing"

	corev1 "k8s.io/api/core/v1"

	"github.com/gruntwork-io/terratest/modules/helm"
)

func TestGrafanaHelmChartTemplate(t *testing.T) {
	// Path to the helm chart we will test
	helmChartGrafanaPath := "../open-electrons-monitoring"

	// Setup the args. For this test, we will set the following input values:
	// - image=grafana:latest
	options := &helm.Options{
		SetValues: map[string]string{"image": "grafana:latest"},
	}

	// Run RenderTemplate to render the template and capture the output.
	output := helm.RenderTemplate(t, options, helmChartGrafanaPath, "pod", []string{"../../templates/grafana/grafana-deployment.yml"})

	// Now we use kubernetes/client-go library to render the template output into the Pod struct. This will
	// ensure the Pod resource is rendered correctly.
	var pod corev1.Pod
	helm.UnmarshalK8SYaml(t, output, &pod)

	// Finally, we verify the pod spec is set to the expected container image value
	expectedContainerImage := "grafana:latest"
	podContainers := pod.Spec.Containers
	if podContainers[0].Image != expectedContainerImage {
		t.Fatalf("Rendered container image (%s) is not expected (%s)", podContainers[0].Image, expectedContainerImage)
	}
}