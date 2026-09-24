// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mwaaserverlessworkflow


type MwaaserverlessWorkflowCode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mwaaserverless_workflow#s3_location MwaaserverlessWorkflow#s3_location}.
	S3Location *MwaaserverlessWorkflowCodeS3Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}

