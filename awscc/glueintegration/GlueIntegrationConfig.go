// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueintegration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueIntegrationConfig struct {
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
	// The name of the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/glue_integration#integration_name GlueIntegration#integration_name}
	IntegrationName *string `field:"required" json:"integrationName" yaml:"integrationName"`
	// The Amazon Resource Name (ARN) of the database to use as the source for replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/glue_integration#source_arn GlueIntegration#source_arn}
	SourceArn *string `field:"required" json:"sourceArn" yaml:"sourceArn"`
	// The Amazon Resource Name (ARN) of the Glue data warehouse to use as the target for replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/glue_integration#target_arn GlueIntegration#target_arn}
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// An optional set of non-secret key value pairs that contains additional contextual information about the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/glue_integration#additional_encryption_context GlueIntegration#additional_encryption_context}
	AdditionalEncryptionContext *map[string]*string `field:"optional" json:"additionalEncryptionContext" yaml:"additionalEncryptionContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/glue_integration#data_filter GlueIntegration#data_filter}.
	DataFilter *string `field:"optional" json:"dataFilter" yaml:"dataFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/glue_integration#description GlueIntegration#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The configuration settings for the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/glue_integration#integration_config GlueIntegration#integration_config}
	IntegrationConfig *GlueIntegrationIntegrationConfig `field:"optional" json:"integrationConfig" yaml:"integrationConfig"`
	// An KMS key identifier for the key to use to encrypt the integration.
	//
	// If you don't specify an encryption key, the default AWS owned KMS key is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/glue_integration#kms_key_id GlueIntegration#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/glue_integration#tags GlueIntegration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

