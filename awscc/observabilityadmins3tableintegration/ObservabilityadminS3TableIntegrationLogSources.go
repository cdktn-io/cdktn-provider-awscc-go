// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmins3tableintegration


type ObservabilityadminS3TableIntegrationLogSources struct {
	// The ID of the CloudWatch Logs data source association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/observabilityadmin_s3_table_integration#identifier ObservabilityadminS3TableIntegration#identifier}
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// The name of the CloudWatch Logs data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/observabilityadmin_s3_table_integration#name ObservabilityadminS3TableIntegration#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of the CloudWatch Logs data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/observabilityadmin_s3_table_integration#type ObservabilityadminS3TableIntegration#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

