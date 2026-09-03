// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterInstanceGroupsAutoPatchConfig struct {
	// The configuration to use when updating the AMI versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_cluster#deployment_config SagemakerCluster#deployment_config}
	DeploymentConfig *SagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
	// The patching strategy that determines when and how instances are patched.
	//
	// WhenIdle patches instances as they become idle. WhenAllIdle patches all instances when they are all idle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_cluster#patching_strategy SagemakerCluster#patching_strategy}
	PatchingStrategy *string `field:"optional" json:"patchingStrategy" yaml:"patchingStrategy"`
	// The schedule configuration for automatic patching.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_cluster#patch_schedule SagemakerCluster#patch_schedule}
	PatchSchedule *SagemakerClusterInstanceGroupsAutoPatchConfigPatchSchedule `field:"optional" json:"patchSchedule" yaml:"patchSchedule"`
}

