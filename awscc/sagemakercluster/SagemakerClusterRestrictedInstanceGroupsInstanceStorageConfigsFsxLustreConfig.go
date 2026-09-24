// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterRestrictedInstanceGroupsInstanceStorageConfigsFsxLustreConfig struct {
	// The DNS name of the FSx for Lustre file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_cluster#dns_name SagemakerCluster#dns_name}
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// The mount name of the FSx for Lustre file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_cluster#mount_name SagemakerCluster#mount_name}
	MountName *string `field:"optional" json:"mountName" yaml:"mountName"`
	// The mount path for the FSx for Lustre file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_cluster#mount_path SagemakerCluster#mount_path}
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

