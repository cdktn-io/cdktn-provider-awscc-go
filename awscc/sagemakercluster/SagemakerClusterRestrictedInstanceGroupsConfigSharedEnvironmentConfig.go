// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfig struct {
	// Configuration settings for an Amazon FSx for Lustre file system to be used with the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#fsx_lustre_config SagemakerCluster#fsx_lustre_config}
	FsxLustreConfig *SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfig `field:"optional" json:"fsxLustreConfig" yaml:"fsxLustreConfig"`
	// The deletion policy for the shared FSx Lustre file system.
	//
	// Keep retains the FSx when RIGs are deleted. DeleteIfNotUsed deletes the FSx when no RIGs reference it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#fsx_lustre_deletion_policy SagemakerCluster#fsx_lustre_deletion_policy}
	FsxLustreDeletionPolicy *string `field:"optional" json:"fsxLustreDeletionPolicy" yaml:"fsxLustreDeletionPolicy"`
}

