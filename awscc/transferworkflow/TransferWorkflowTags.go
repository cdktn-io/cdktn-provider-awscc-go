// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferworkflow


type TransferWorkflowTags struct {
	// The name assigned to the tag that you create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/transfer_workflow#key TransferWorkflow#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Contains one or more values that you assigned to the key name you create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/transfer_workflow#value TransferWorkflow#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

