// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferworkflow


type TransferWorkflowOnExceptionSteps struct {
	// Details for a step that performs a file copy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/transfer_workflow#copy_step_details TransferWorkflow#copy_step_details}
	CopyStepDetails *TransferWorkflowOnExceptionStepsCopyStepDetails `field:"optional" json:"copyStepDetails" yaml:"copyStepDetails"`
	// Details for a step that invokes a lambda function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/transfer_workflow#custom_step_details TransferWorkflow#custom_step_details}
	CustomStepDetails *TransferWorkflowOnExceptionStepsCustomStepDetails `field:"optional" json:"customStepDetails" yaml:"customStepDetails"`
	// Details for a step that performs a file decryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/transfer_workflow#decrypt_step_details TransferWorkflow#decrypt_step_details}
	DecryptStepDetails *TransferWorkflowOnExceptionStepsDecryptStepDetails `field:"optional" json:"decryptStepDetails" yaml:"decryptStepDetails"`
	// Details for a step that deletes the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/transfer_workflow#delete_step_details TransferWorkflow#delete_step_details}
	DeleteStepDetails *TransferWorkflowOnExceptionStepsDeleteStepDetails `field:"optional" json:"deleteStepDetails" yaml:"deleteStepDetails"`
	// Details for a step that creates one or more tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/transfer_workflow#tag_step_details TransferWorkflow#tag_step_details}
	TagStepDetails *TransferWorkflowOnExceptionStepsTagStepDetails `field:"optional" json:"tagStepDetails" yaml:"tagStepDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/transfer_workflow#type TransferWorkflow#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

