// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferworkflow


type TransferWorkflowStepsCopyStepDetailsDestinationFileLocationS3FileLocation struct {
	// Specifies the S3 bucket that contains the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/transfer_workflow#bucket TransferWorkflow#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The name assigned to the file when it was created in S3.
	//
	// You use the object key to retrieve the object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/transfer_workflow#key TransferWorkflow#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
}

