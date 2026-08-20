// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememory


type BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSelfManagedConfigurationInvocationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_memory#payload_delivery_bucket_name BedrockagentcoreMemory#payload_delivery_bucket_name}.
	PayloadDeliveryBucketName *string `field:"optional" json:"payloadDeliveryBucketName" yaml:"payloadDeliveryBucketName"`
	// ARN format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_memory#topic_arn BedrockagentcoreMemory#topic_arn}
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

