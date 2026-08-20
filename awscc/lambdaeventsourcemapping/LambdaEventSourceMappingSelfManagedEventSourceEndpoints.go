// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdaeventsourcemapping


type LambdaEventSourceMappingSelfManagedEventSourceEndpoints struct {
	// The list of bootstrap servers for your Kafka brokers in the following format: ``"KafkaBootstrapServers": ["abc.xyz.com:xxxx","abc2.xyz.com:xxxx"]``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#kafka_bootstrap_servers LambdaEventSourceMapping#kafka_bootstrap_servers}
	KafkaBootstrapServers *[]*string `field:"optional" json:"kafkaBootstrapServers" yaml:"kafkaBootstrapServers"`
}

