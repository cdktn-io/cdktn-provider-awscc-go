// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmins3tableintegration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityadminS3TableIntegrationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Encryption configuration for the S3 Table Integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_s3_table_integration#encryption ObservabilityadminS3TableIntegration#encryption}
	Encryption *ObservabilityadminS3TableIntegrationEncryption `field:"required" json:"encryption" yaml:"encryption"`
	// The ARN of the role used to access the S3 Table Integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_s3_table_integration#role_arn ObservabilityadminS3TableIntegration#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The CloudWatch Logs data sources to associate with the S3 Table Integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_s3_table_integration#log_sources ObservabilityadminS3TableIntegration#log_sources}
	LogSources interface{} `field:"optional" json:"logSources" yaml:"logSources"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_s3_table_integration#tags ObservabilityadminS3TableIntegration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

