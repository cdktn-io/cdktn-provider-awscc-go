// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilesintegration


type CustomerprofilesIntegrationFlowDefinitionTasks struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_integration#connector_operator CustomerprofilesIntegration#connector_operator}.
	ConnectorOperator *CustomerprofilesIntegrationFlowDefinitionTasksConnectorOperator `field:"optional" json:"connectorOperator" yaml:"connectorOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_integration#destination_field CustomerprofilesIntegration#destination_field}.
	DestinationField *string `field:"optional" json:"destinationField" yaml:"destinationField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_integration#source_fields CustomerprofilesIntegration#source_fields}.
	SourceFields *[]*string `field:"optional" json:"sourceFields" yaml:"sourceFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_integration#task_properties CustomerprofilesIntegration#task_properties}.
	TaskProperties interface{} `field:"optional" json:"taskProperties" yaml:"taskProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_integration#task_type CustomerprofilesIntegration#task_type}.
	TaskType *string `field:"optional" json:"taskType" yaml:"taskType"`
}

