package test

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/random"
)

// This file contains examples of how to use terratest to test helm charts by deploying the chart and verifying the
// deployment by hitting the service endpoint.
func TestSelfSignedAdmissionController(t *testing.T) {
	t.Parallel()

	// Path to the helm chart
	helmChartPath, err := filepath.Abs("../../charts/opa-sidecar-admission-controller")
	require.NoError(t, err)

	// Define unique test namespace.
	namespaceName := fmt.Sprintf("opa-sidecar-admission-controller-%s", strings.ToLower(random.UniqueId()))
	
	// Setup default kubectl config and context.
	// - HOME/.kube/config for the kubectl config file
	// - Current context of the kubectl config file
	kubectlOptions := k8s.NewKubectlOptions("", "", namespaceName)

	k8s.CreateNamespace(t, kubectlOptions, namespaceName)
	defer k8s.DeleteNamespace(t, kubectlOptions, namespaceName)

	// Setup the helm cli args
	// - values= self-signed ac option
	options := &helm.Options{
		KubectlOptions: kubectlOptions,
		ValuesFiles: []string{
			"../../charts/opa-sidecar-admission-controller/values.yaml",
		},
	}

	// Generate a unique release name
	releaseName := fmt.Sprintf(
		"opa-sidecar-admission-controller-%s",
		strings.ToLower(random.UniqueId()),
	)
	defer helm.Delete(t, options, releaseName, true)
	// Deploy the chart using `helm install`. Note that we use the version without `E`, since we want to assert the
	// install succeeds without any errors.
	helm.Install(t, options, helmChartPath, releaseName)

	// Verify the deployment.
	k8s.WaitUntilSecretAvailable(t, kubectlOptions, fmt.Sprintf("%s-certificate", releaseName), 5, time.Duration(3) * time.Second)
	k8s.RunKubectl(t, kubectlOptions, "get", "mutatingwebhookconfiguration", fmt.Sprintf("%s-webhook", releaseName))
	
}