// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememory


type BedrockagentcoreMemoryStreamDeliveryResourcesResourcesKinesis struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_memory#content_configurations BedrockagentcoreMemory#content_configurations}.
	ContentConfigurations interface{} `field:"optional" json:"contentConfigurations" yaml:"contentConfigurations"`
	// ARN format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_memory#data_stream_arn BedrockagentcoreMemory#data_stream_arn}
	DataStreamArn *string `field:"optional" json:"dataStreamArn" yaml:"dataStreamArn"`
}

