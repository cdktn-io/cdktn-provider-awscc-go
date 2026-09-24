// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdaeventsourcemapping


type LambdaEventSourceMappingAmazonManagedKafkaEventSourceConfigSchemaRegistryConfigAccessConfigs struct {
	// The type of authentication Lambda uses to access your schema registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_event_source_mapping#type LambdaEventSourceMapping#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The URI of the secret (Secrets Manager secret ARN) to authenticate with your schema registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_event_source_mapping#uri LambdaEventSourceMapping#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

