// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowDestinationFlowConfigListStruct struct {
	// Destination connector type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_flow#connector_type AppflowFlow#connector_type}
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// Destination connector details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_flow#destination_connector_properties AppflowFlow#destination_connector_properties}
	DestinationConnectorProperties *AppflowFlowDestinationFlowConfigListDestinationConnectorProperties `field:"required" json:"destinationConnectorProperties" yaml:"destinationConnectorProperties"`
	// The API version that the destination connector uses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_flow#api_version AppflowFlow#api_version}
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Name of destination connector profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_flow#connector_profile_name AppflowFlow#connector_profile_name}
	ConnectorProfileName *string `field:"optional" json:"connectorProfileName" yaml:"connectorProfileName"`
}

