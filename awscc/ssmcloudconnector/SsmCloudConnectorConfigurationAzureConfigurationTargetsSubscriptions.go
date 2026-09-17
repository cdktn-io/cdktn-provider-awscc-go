// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmcloudconnector


type SsmCloudConnectorConfigurationAzureConfigurationTargetsSubscriptions struct {
	// The display name of the Azure subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_cloud_connector#display_name SsmCloudConnector#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The Azure subscription ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_cloud_connector#id SsmCloudConnector#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

