// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package omicsworkflow


type OmicsWorkflowParameterTemplate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/omics_workflow#description OmicsWorkflow#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/omics_workflow#optional OmicsWorkflow#optional}.
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

