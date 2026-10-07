// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdaeventsourcemapping


type LambdaEventSourceMappingSelfManagedKafkaEventSourceConfig struct {
	// The identifier for the Kafka consumer group to join.
	//
	// The consumer group ID must be unique among all your Kafka event sources. After creating a Kafka event source mapping with the consumer group ID specified, you cannot update this value. For more information, see [Customizable consumer group ID](https://docs.aws.amazon.com/lambda/latest/dg/with-kafka-process.html#services-smaa-topic-add).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/lambda_event_source_mapping#consumer_group_id LambdaEventSourceMapping#consumer_group_id}
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
	// The mode that determines how Lambda reads from a Kafka topic.
	//
	// Use Stream for ordered processing or Queue for higher throughput when ordering is not required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/lambda_event_source_mapping#consumption_mode LambdaEventSourceMapping#consumption_mode}
	ConsumptionMode *string `field:"optional" json:"consumptionMode" yaml:"consumptionMode"`
	// Specific configuration settings for a Kafka schema registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/lambda_event_source_mapping#schema_registry_config LambdaEventSourceMapping#schema_registry_config}
	SchemaRegistryConfig *LambdaEventSourceMappingSelfManagedKafkaEventSourceConfigSchemaRegistryConfig `field:"optional" json:"schemaRegistryConfig" yaml:"schemaRegistryConfig"`
}

