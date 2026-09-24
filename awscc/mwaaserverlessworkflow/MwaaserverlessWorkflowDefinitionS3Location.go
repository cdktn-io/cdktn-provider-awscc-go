// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mwaaserverlessworkflow


type MwaaserverlessWorkflowDefinitionS3Location struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mwaaserverless_workflow#bucket MwaaserverlessWorkflow#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mwaaserverless_workflow#object_key MwaaserverlessWorkflow#object_key}.
	ObjectKey *string `field:"required" json:"objectKey" yaml:"objectKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mwaaserverless_workflow#version_id MwaaserverlessWorkflow#version_id}.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

