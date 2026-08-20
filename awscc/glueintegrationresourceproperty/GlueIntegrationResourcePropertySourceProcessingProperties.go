// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueintegrationresourceproperty


type GlueIntegrationResourcePropertySourceProcessingProperties struct {
	// The IAM role to access the Glue connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/glue_integration_resource_property#role_arn GlueIntegrationResourceProperty#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

