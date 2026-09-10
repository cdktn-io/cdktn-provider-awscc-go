// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagent


type BedrockAgentActionGroupsApiSchema struct {
	// String OpenAPI Payload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_agent#payload BedrockAgent#payload}
	Payload *string `field:"optional" json:"payload" yaml:"payload"`
	// The identifier for the S3 resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_agent#s3 BedrockAgent#s3}
	S3 *BedrockAgentActionGroupsApiSchemaS3 `field:"optional" json:"s3" yaml:"s3"`
}

