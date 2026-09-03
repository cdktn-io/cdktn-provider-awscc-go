// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnector struct {
	// A map for properties for custom connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_flow#custom_properties AppflowFlow#custom_properties}
	CustomProperties *map[string]*string `field:"optional" json:"customProperties" yaml:"customProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_flow#data_transfer_api AppflowFlow#data_transfer_api}.
	DataTransferApi *AppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnectorDataTransferApi `field:"optional" json:"dataTransferApi" yaml:"dataTransferApi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_flow#entity_name AppflowFlow#entity_name}.
	EntityName *string `field:"optional" json:"entityName" yaml:"entityName"`
}

