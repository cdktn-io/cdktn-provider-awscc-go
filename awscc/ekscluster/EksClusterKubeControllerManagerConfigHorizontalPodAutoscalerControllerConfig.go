// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterKubeControllerManagerConfigHorizontalPodAutoscalerControllerConfig struct {
	// The interval between each sync of the horizontal pod autoscaler (e.g., 15s, 1m).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/eks_cluster#horizontal_pod_autoscaler_sync_period EksCluster#horizontal_pod_autoscaler_sync_period}
	HorizontalPodAutoscalerSyncPeriod *string `field:"optional" json:"horizontalPodAutoscalerSyncPeriod" yaml:"horizontalPodAutoscalerSyncPeriod"`
}

