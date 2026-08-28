// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securitylakedatalake


type SecuritylakeDataLakeLifecycleConfigurationTransitions struct {
	// Number of days before data transitions to a different S3 Storage Class in the Amazon Security Lake object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securitylake_data_lake#days SecuritylakeDataLake#days}
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// The range of storage classes that you can choose from based on the data access, resiliency, and cost requirements of your workloads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securitylake_data_lake#storage_class SecuritylakeDataLake#storage_class}
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

