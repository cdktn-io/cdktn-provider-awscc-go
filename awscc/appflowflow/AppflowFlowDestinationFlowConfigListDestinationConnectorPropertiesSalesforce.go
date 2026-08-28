// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSalesforce struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appflow_flow#data_transfer_api AppflowFlow#data_transfer_api}.
	DataTransferApi *string `field:"optional" json:"dataTransferApi" yaml:"dataTransferApi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appflow_flow#error_handling_config AppflowFlow#error_handling_config}.
	ErrorHandlingConfig *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSalesforceErrorHandlingConfig `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
	// List of fields used as ID when performing a write operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appflow_flow#id_field_names AppflowFlow#id_field_names}
	IdFieldNames *[]*string `field:"optional" json:"idFieldNames" yaml:"idFieldNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appflow_flow#object AppflowFlow#object}.
	Object *string `field:"optional" json:"object" yaml:"object"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appflow_flow#write_operation_type AppflowFlow#write_operation_type}.
	WriteOperationType *string `field:"optional" json:"writeOperationType" yaml:"writeOperationType"`
}

