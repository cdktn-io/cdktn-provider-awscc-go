// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinepipeline


type CodepipelinePipelineArtifactStore struct {
	// Represents information about the key used to encrypt data in the artifact store, such as an AWS Key Management Service (AWS KMS) key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codepipeline_pipeline#encryption_key CodepipelinePipeline#encryption_key}
	EncryptionKey *CodepipelinePipelineArtifactStoreEncryptionKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The S3 bucket used for storing the artifacts for a pipeline.
	//
	// You can specify the name of an S3 bucket but not a folder in the bucket. A folder to contain the pipeline artifacts is created for you based on the name of the pipeline. You can use any S3 bucket in the same AWS Region as the pipeline to store your pipeline artifacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codepipeline_pipeline#location CodepipelinePipeline#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The type of the artifact store, such as S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codepipeline_pipeline#type CodepipelinePipeline#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

