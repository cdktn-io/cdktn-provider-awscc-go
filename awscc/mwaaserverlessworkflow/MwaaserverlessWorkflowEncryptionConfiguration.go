// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mwaaserverlessworkflow


type MwaaserverlessWorkflowEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mwaaserverless_workflow#kms_key_id MwaaserverlessWorkflow#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mwaaserverless_workflow#type MwaaserverlessWorkflow#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

