// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdaeventsourcemapping


type LambdaEventSourceMappingAmazonManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigs struct {
	// The attributes you want your schema registry to validate and filter for.
	//
	// If you selected ``JSON`` as the ``EventRecordFormat``, Lambda also deserializes the selected message attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_event_source_mapping#attribute LambdaEventSourceMapping#attribute}
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
}

