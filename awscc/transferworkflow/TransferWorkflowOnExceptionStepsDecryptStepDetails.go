// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferworkflow


type TransferWorkflowOnExceptionStepsDecryptStepDetails struct {
	// Specifies the location for the file being decrypted. Only applicable for the Decrypt type of workflow steps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/transfer_workflow#destination_file_location TransferWorkflow#destination_file_location}
	DestinationFileLocation *TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocation `field:"optional" json:"destinationFileLocation" yaml:"destinationFileLocation"`
	// The name of the step, used as an identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/transfer_workflow#name TransferWorkflow#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A flag that indicates whether or not to overwrite an existing file of the same name.
	//
	// The default is FALSE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/transfer_workflow#overwrite_existing TransferWorkflow#overwrite_existing}
	OverwriteExisting *string `field:"optional" json:"overwriteExisting" yaml:"overwriteExisting"`
	// Specifies which file to use as input to the workflow step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/transfer_workflow#source_file_location TransferWorkflow#source_file_location}
	SourceFileLocation *string `field:"optional" json:"sourceFileLocation" yaml:"sourceFileLocation"`
	// Specifies which encryption method to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/transfer_workflow#type TransferWorkflow#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

