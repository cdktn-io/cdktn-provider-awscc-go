// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmcloudconnector


type SsmCloudConnectorConfigurationAzureConfigurationTargets struct {
	// List of Azure subscriptions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssm_cloud_connector#subscriptions SsmCloudConnector#subscriptions}
	Subscriptions interface{} `field:"optional" json:"subscriptions" yaml:"subscriptions"`
}

