// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterRestrictedInstanceGroupsEnvironmentConfig struct {
	// Configuration settings for an Amazon FSx for Lustre file system to be used with the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#fsx_lustre_config SagemakerCluster#fsx_lustre_config}
	FsxLustreConfig *SagemakerClusterRestrictedInstanceGroupsEnvironmentConfigFsxLustreConfig `field:"optional" json:"fsxLustreConfig" yaml:"fsxLustreConfig"`
}

