// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeSourceParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pipes_pipe#active_mq_broker_parameters PipesPipe#active_mq_broker_parameters}.
	ActiveMqBrokerParameters *PipesPipeSourceParametersActiveMqBrokerParameters `field:"optional" json:"activeMqBrokerParameters" yaml:"activeMqBrokerParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pipes_pipe#dynamo_db_stream_parameters PipesPipe#dynamo_db_stream_parameters}.
	DynamoDbStreamParameters *PipesPipeSourceParametersDynamoDbStreamParameters `field:"optional" json:"dynamoDbStreamParameters" yaml:"dynamoDbStreamParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pipes_pipe#filter_criteria PipesPipe#filter_criteria}.
	FilterCriteria *PipesPipeSourceParametersFilterCriteria `field:"optional" json:"filterCriteria" yaml:"filterCriteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pipes_pipe#kinesis_stream_parameters PipesPipe#kinesis_stream_parameters}.
	KinesisStreamParameters *PipesPipeSourceParametersKinesisStreamParameters `field:"optional" json:"kinesisStreamParameters" yaml:"kinesisStreamParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pipes_pipe#managed_streaming_kafka_parameters PipesPipe#managed_streaming_kafka_parameters}.
	ManagedStreamingKafkaParameters *PipesPipeSourceParametersManagedStreamingKafkaParameters `field:"optional" json:"managedStreamingKafkaParameters" yaml:"managedStreamingKafkaParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pipes_pipe#rabbit_mq_broker_parameters PipesPipe#rabbit_mq_broker_parameters}.
	RabbitMqBrokerParameters *PipesPipeSourceParametersRabbitMqBrokerParameters `field:"optional" json:"rabbitMqBrokerParameters" yaml:"rabbitMqBrokerParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pipes_pipe#self_managed_kafka_parameters PipesPipe#self_managed_kafka_parameters}.
	SelfManagedKafkaParameters *PipesPipeSourceParametersSelfManagedKafkaParameters `field:"optional" json:"selfManagedKafkaParameters" yaml:"selfManagedKafkaParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pipes_pipe#sqs_queue_parameters PipesPipe#sqs_queue_parameters}.
	SqsQueueParameters *PipesPipeSourceParametersSqsQueueParameters `field:"optional" json:"sqsQueueParameters" yaml:"sqsQueueParameters"`
}

