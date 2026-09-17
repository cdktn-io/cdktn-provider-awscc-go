// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferworkflow


type TransferWorkflowStepsDecryptStepDetailsDestinationFileLocation struct {
	// Specifies the details for an EFS file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/transfer_workflow#efs_file_location TransferWorkflow#efs_file_location}
	EfsFileLocation *TransferWorkflowStepsDecryptStepDetailsDestinationFileLocationEfsFileLocation `field:"optional" json:"efsFileLocation" yaml:"efsFileLocation"`
	// Specifies the details for a S3 file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/transfer_workflow#s3_file_location TransferWorkflow#s3_file_location}
	S3FileLocation *TransferWorkflowStepsDecryptStepDetailsDestinationFileLocationS3FileLocation `field:"optional" json:"s3FileLocation" yaml:"s3FileLocation"`
}

