// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercontext


type SagemakerContextSource struct {
	// The URI of the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_context#source_uri SagemakerContext#source_uri}
	SourceUri *string `field:"required" json:"sourceUri" yaml:"sourceUri"`
	// The ID of the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_context#source_id SagemakerContext#source_id}
	SourceId *string `field:"optional" json:"sourceId" yaml:"sourceId"`
	// The type of the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_context#source_type SagemakerContext#source_type}
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}

