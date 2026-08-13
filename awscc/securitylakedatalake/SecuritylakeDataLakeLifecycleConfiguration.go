// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securitylakedatalake


type SecuritylakeDataLakeLifecycleConfiguration struct {
	// Provides data expiration details of Amazon Security Lake object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securitylake_data_lake#expiration SecuritylakeDataLake#expiration}
	Expiration *SecuritylakeDataLakeLifecycleConfigurationExpiration `field:"optional" json:"expiration" yaml:"expiration"`
	// Provides data storage transition details of Amazon Security Lake object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securitylake_data_lake#transitions SecuritylakeDataLake#transitions}
	Transitions interface{} `field:"optional" json:"transitions" yaml:"transitions"`
}

