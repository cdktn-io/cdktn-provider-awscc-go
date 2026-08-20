// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdaeventsourcemapping


type LambdaEventSourceMappingSelfManagedEventSource struct {
	// The list of bootstrap servers for your Kafka brokers in the following format: ``"KafkaBootstrapServers": ["abc.xyz.com:xxxx","abc2.xyz.com:xxxx"]``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#endpoints LambdaEventSourceMapping#endpoints}
	Endpoints *LambdaEventSourceMappingSelfManagedEventSourceEndpoints `field:"optional" json:"endpoints" yaml:"endpoints"`
}

