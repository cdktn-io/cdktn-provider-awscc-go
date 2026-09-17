// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdaeventsourcemapping


type LambdaEventSourceMappingSelfManagedKafkaEventSourceConfigSchemaRegistryConfig struct {
	// An array of access configuration objects that tell Lambda how to authenticate with your schema registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_event_source_mapping#access_configs LambdaEventSourceMapping#access_configs}
	AccessConfigs interface{} `field:"optional" json:"accessConfigs" yaml:"accessConfigs"`
	// The record format that Lambda delivers to your function after schema validation.
	//
	// +  Choose ``JSON`` to have Lambda deliver the record to your function as a standard JSON object.
	//   +  Choose ``SOURCE`` to have Lambda deliver the record to your function in its original source format. Lambda removes all schema metadata, such as the schema ID, before sending the record to your function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_event_source_mapping#event_record_format LambdaEventSourceMapping#event_record_format}
	EventRecordFormat *string `field:"optional" json:"eventRecordFormat" yaml:"eventRecordFormat"`
	// The URI for your schema registry.
	//
	// The correct URI format depends on the type of schema registry you're using.
	//   +  For GLU schema registries, use the ARN of the registry.
	//   +  For Confluent schema registries, use the URL of the registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_event_source_mapping#schema_registry_uri LambdaEventSourceMapping#schema_registry_uri}
	SchemaRegistryUri *string `field:"optional" json:"schemaRegistryUri" yaml:"schemaRegistryUri"`
	// An array of schema validation configuration objects, which tell Lambda the message attributes you want to validate and filter using your schema registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_event_source_mapping#schema_validation_configs LambdaEventSourceMapping#schema_validation_configs}
	SchemaValidationConfigs interface{} `field:"optional" json:"schemaValidationConfigs" yaml:"schemaValidationConfigs"`
}

