// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeraction


type SagemakerActionSource struct {
	// The URI of the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_action#source_uri SagemakerAction#source_uri}
	SourceUri *string `field:"required" json:"sourceUri" yaml:"sourceUri"`
	// The ID of the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_action#source_id SagemakerAction#source_id}
	SourceId *string `field:"optional" json:"sourceId" yaml:"sourceId"`
	// The type of the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_action#source_type SagemakerAction#source_type}
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}

