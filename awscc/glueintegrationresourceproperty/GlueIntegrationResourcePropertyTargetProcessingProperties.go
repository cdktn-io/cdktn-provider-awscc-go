// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueintegrationresourceproperty


type GlueIntegrationResourcePropertyTargetProcessingProperties struct {
	// The Glue network connection to configure the Glue job running in the customer VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/glue_integration_resource_property#connection_name GlueIntegrationResourceProperty#connection_name}
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// The ARN of an Eventbridge event bus to receive the integration status notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/glue_integration_resource_property#event_bus_arn GlueIntegrationResourceProperty#event_bus_arn}
	EventBusArn *string `field:"optional" json:"eventBusArn" yaml:"eventBusArn"`
	// The ARN of the KMS key used for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/glue_integration_resource_property#kms_arn GlueIntegrationResourceProperty#kms_arn}
	KmsArn *string `field:"optional" json:"kmsArn" yaml:"kmsArn"`
	// The IAM role to access the Glue database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/glue_integration_resource_property#role_arn GlueIntegrationResourceProperty#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

