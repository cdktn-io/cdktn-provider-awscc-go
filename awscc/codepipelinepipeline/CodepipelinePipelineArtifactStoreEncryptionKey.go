// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline


type CodepipelinePipelineArtifactStoreEncryptionKey struct {
	// The ID used to identify the key.
	//
	// For an AWS KMS key, you can use the key ID, the key ARN, or the alias ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/codepipeline_pipeline#id CodepipelinePipeline#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The type of encryption key, such as an AWS KMS key.
	//
	// When creating or updating a pipeline, the value must be set to 'KMS'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/codepipeline_pipeline#type CodepipelinePipeline#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

