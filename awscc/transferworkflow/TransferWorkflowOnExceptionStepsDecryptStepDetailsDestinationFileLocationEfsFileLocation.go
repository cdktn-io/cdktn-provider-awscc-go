// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferworkflow


type TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationEfsFileLocation struct {
	// Specifies the EFS filesystem that contains the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/transfer_workflow#file_system_id TransferWorkflow#file_system_id}
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// The name assigned to the file when it was created in EFS.
	//
	// You use the object path to retrieve the object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/transfer_workflow#path TransferWorkflow#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

