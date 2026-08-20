// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowSourceFlowConfig struct {
	// Type of source connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appflow_flow#connector_type AppflowFlow#connector_type}
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// Source connector details required to query a connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appflow_flow#source_connector_properties AppflowFlow#source_connector_properties}
	SourceConnectorProperties *AppflowFlowSourceFlowConfigSourceConnectorProperties `field:"required" json:"sourceConnectorProperties" yaml:"sourceConnectorProperties"`
	// The API version that the destination connector uses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appflow_flow#api_version AppflowFlow#api_version}
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Name of source connector profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appflow_flow#connector_profile_name AppflowFlow#connector_profile_name}
	ConnectorProfileName *string `field:"optional" json:"connectorProfileName" yaml:"connectorProfileName"`
	// Configuration for scheduled incremental data pull.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appflow_flow#incremental_pull_config AppflowFlow#incremental_pull_config}
	IncrementalPullConfig *AppflowFlowSourceFlowConfigIncrementalPullConfig `field:"optional" json:"incrementalPullConfig" yaml:"incrementalPullConfig"`
}

