// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncchannelnamespace


type AppsyncChannelNamespaceHandlerConfigsOnPublishIntegration struct {
	// Data source to invoke for this integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appsync_channel_namespace#data_source_name AppsyncChannelNamespace#data_source_name}
	DataSourceName *string `field:"optional" json:"dataSourceName" yaml:"dataSourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appsync_channel_namespace#lambda_config AppsyncChannelNamespace#lambda_config}.
	LambdaConfig *AppsyncChannelNamespaceHandlerConfigsOnPublishIntegrationLambdaConfig `field:"optional" json:"lambdaConfig" yaml:"lambdaConfig"`
}

