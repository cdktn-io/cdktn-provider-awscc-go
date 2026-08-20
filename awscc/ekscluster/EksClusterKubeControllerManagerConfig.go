// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterKubeControllerManagerConfig struct {
	// The horizontal pod autoscaler controller configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#horizontal_pod_autoscaler_controller_config EksCluster#horizontal_pod_autoscaler_controller_config}
	HorizontalPodAutoscalerControllerConfig *EksClusterKubeControllerManagerConfigHorizontalPodAutoscalerControllerConfig `field:"optional" json:"horizontalPodAutoscalerControllerConfig" yaml:"horizontalPodAutoscalerControllerConfig"`
}

