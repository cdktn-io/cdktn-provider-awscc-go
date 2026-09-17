// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterRestrictedInstanceGroupsConfig struct {
	// The shared environment configuration for restricted instance groups that use cluster-level shared FSx Lustre storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#shared_environment_config SagemakerCluster#shared_environment_config}
	SharedEnvironmentConfig *SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfig `field:"optional" json:"sharedEnvironmentConfig" yaml:"sharedEnvironmentConfig"`
}

