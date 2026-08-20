// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationgeneratedtemplate


type CloudformationGeneratedTemplateTemplateConfiguration struct {
	// The DeletionPolicy assigned to resources in the generated template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudformation_generated_template#deletion_policy CloudformationGeneratedTemplate#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// The UpdateReplacePolicy assigned to resources in the generated template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudformation_generated_template#update_replace_policy CloudformationGeneratedTemplate#update_replace_policy}
	UpdateReplacePolicy *string `field:"optional" json:"updateReplacePolicy" yaml:"updateReplacePolicy"`
}

