// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterRestrictedInstanceGroupsInstanceStorageConfigs struct {
	// Defines the configuration for attaching additional Amazon Elastic Block Store (EBS) volumes to the instances in the SageMaker HyperPod cluster instance group.
	//
	// The additional EBS volume is attached to each instance within the SageMaker HyperPod cluster instance group and mounted to /opt/sagemaker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#ebs_volume_config SagemakerCluster#ebs_volume_config}
	EbsVolumeConfig *SagemakerClusterRestrictedInstanceGroupsInstanceStorageConfigsEbsVolumeConfig `field:"optional" json:"ebsVolumeConfig" yaml:"ebsVolumeConfig"`
	// Configuration for mounting an Amazon FSx Lustre file system to the instances in the SageMaker HyperPod cluster instance group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#fsx_lustre_config SagemakerCluster#fsx_lustre_config}
	FsxLustreConfig *SagemakerClusterRestrictedInstanceGroupsInstanceStorageConfigsFsxLustreConfig `field:"optional" json:"fsxLustreConfig" yaml:"fsxLustreConfig"`
	// Configuration for mounting an Amazon FSx OpenZFS file system to the instances in the SageMaker HyperPod cluster instance group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#fsx_open_zfs_config SagemakerCluster#fsx_open_zfs_config}
	FsxOpenZfsConfig *SagemakerClusterRestrictedInstanceGroupsInstanceStorageConfigsFsxOpenZfsConfig `field:"optional" json:"fsxOpenZfsConfig" yaml:"fsxOpenZfsConfig"`
}

