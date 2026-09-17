// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowDestinationFlowConfigListDestinationConnectorProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#custom_connector AppflowFlow#custom_connector}.
	CustomConnector *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesCustomConnector `field:"optional" json:"customConnector" yaml:"customConnector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#event_bridge AppflowFlow#event_bridge}.
	EventBridge *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesEventBridge `field:"optional" json:"eventBridge" yaml:"eventBridge"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#lookout_metrics AppflowFlow#lookout_metrics}.
	LookoutMetrics *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesLookoutMetrics `field:"optional" json:"lookoutMetrics" yaml:"lookoutMetrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#marketo AppflowFlow#marketo}.
	Marketo *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesMarketo `field:"optional" json:"marketo" yaml:"marketo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#redshift AppflowFlow#redshift}.
	Redshift *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesRedshift `field:"optional" json:"redshift" yaml:"redshift"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#s3 AppflowFlow#s3}.
	S3 *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3 `field:"optional" json:"s3" yaml:"s3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#salesforce AppflowFlow#salesforce}.
	Salesforce *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSalesforce `field:"optional" json:"salesforce" yaml:"salesforce"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#sapo_data AppflowFlow#sapo_data}.
	SapoData *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoData `field:"optional" json:"sapoData" yaml:"sapoData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#snowflake AppflowFlow#snowflake}.
	Snowflake *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflake `field:"optional" json:"snowflake" yaml:"snowflake"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#upsolver AppflowFlow#upsolver}.
	Upsolver *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesUpsolver `field:"optional" json:"upsolver" yaml:"upsolver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#zendesk AppflowFlow#zendesk}.
	Zendesk *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendesk `field:"optional" json:"zendesk" yaml:"zendesk"`
}

