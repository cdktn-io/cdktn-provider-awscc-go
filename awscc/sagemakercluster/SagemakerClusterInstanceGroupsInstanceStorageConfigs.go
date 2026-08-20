// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterInstanceGroupsInstanceStorageConfigs struct {
	// Defines the configuration for attaching additional Amazon Elastic Block Store (EBS) volumes to the instances in the SageMaker HyperPod cluster instance group.
	//
	// The additional EBS volume is attached to each instance within the SageMaker HyperPod cluster instance group and mounted to /opt/sagemaker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_cluster#ebs_volume_config SagemakerCluster#ebs_volume_config}
	EbsVolumeConfig *SagemakerClusterInstanceGroupsInstanceStorageConfigsEbsVolumeConfig `field:"optional" json:"ebsVolumeConfig" yaml:"ebsVolumeConfig"`
}

