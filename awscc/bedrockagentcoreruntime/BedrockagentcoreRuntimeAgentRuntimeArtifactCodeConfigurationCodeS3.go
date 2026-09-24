// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreruntime


type BedrockagentcoreRuntimeAgentRuntimeArtifactCodeConfigurationCodeS3 struct {
	// S3 bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_runtime#bucket BedrockagentcoreRuntime#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// S3 object key prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_runtime#prefix BedrockagentcoreRuntime#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// S3 object version ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_runtime#version_id BedrockagentcoreRuntime#version_id}
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

