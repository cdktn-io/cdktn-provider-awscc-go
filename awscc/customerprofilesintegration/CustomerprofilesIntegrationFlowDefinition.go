// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilesintegration


type CustomerprofilesIntegrationFlowDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/customerprofiles_integration#description CustomerprofilesIntegration#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/customerprofiles_integration#flow_name CustomerprofilesIntegration#flow_name}.
	FlowName *string `field:"optional" json:"flowName" yaml:"flowName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/customerprofiles_integration#kms_arn CustomerprofilesIntegration#kms_arn}.
	KmsArn *string `field:"optional" json:"kmsArn" yaml:"kmsArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/customerprofiles_integration#source_flow_config CustomerprofilesIntegration#source_flow_config}.
	SourceFlowConfig *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfig `field:"optional" json:"sourceFlowConfig" yaml:"sourceFlowConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/customerprofiles_integration#tasks CustomerprofilesIntegration#tasks}.
	Tasks interface{} `field:"optional" json:"tasks" yaml:"tasks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/customerprofiles_integration#trigger_config CustomerprofilesIntegration#trigger_config}.
	TriggerConfig *CustomerprofilesIntegrationFlowDefinitionTriggerConfig `field:"optional" json:"triggerConfig" yaml:"triggerConfig"`
}

