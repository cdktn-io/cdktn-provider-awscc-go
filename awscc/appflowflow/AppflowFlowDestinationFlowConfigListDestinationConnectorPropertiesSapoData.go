// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appflow_flow#error_handling_config AppflowFlow#error_handling_config}.
	ErrorHandlingConfig *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfig `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
	// List of fields used as ID when performing a write operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appflow_flow#id_field_names AppflowFlow#id_field_names}
	IdFieldNames *[]*string `field:"optional" json:"idFieldNames" yaml:"idFieldNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appflow_flow#object_path AppflowFlow#object_path}.
	ObjectPath *string `field:"optional" json:"objectPath" yaml:"objectPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appflow_flow#success_response_handling_config AppflowFlow#success_response_handling_config}.
	SuccessResponseHandlingConfig *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataSuccessResponseHandlingConfig `field:"optional" json:"successResponseHandlingConfig" yaml:"successResponseHandlingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appflow_flow#write_operation_type AppflowFlow#write_operation_type}.
	WriteOperationType *string `field:"optional" json:"writeOperationType" yaml:"writeOperationType"`
}

