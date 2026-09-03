// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterInstanceGroupsInstanceStorageConfigsFsxOpenZfsConfig struct {
	// The DNS name of the FSx for OpenZFS file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_cluster#dns_name SagemakerCluster#dns_name}
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// The mount path for the FSx for OpenZFS file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_cluster#mount_path SagemakerCluster#mount_path}
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

