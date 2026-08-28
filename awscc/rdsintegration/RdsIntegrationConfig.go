// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsintegration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RdsIntegrationConfig struct {
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
	// The Amazon Resource Name (ARN) of the database to use as the source for replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_integration#source_arn RdsIntegration#source_arn}
	SourceArn *string `field:"required" json:"sourceArn" yaml:"sourceArn"`
	// The ARN of the Redshift data warehouse to use as the target for replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_integration#target_arn RdsIntegration#target_arn}
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// An optional set of non-secret key?value pairs that contains additional contextual information about the data. For more information, see [Encryption context](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context) in the *Key Management Service Developer Guide*.  You can only include this parameter if you specify the ``KMSKeyId`` parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_integration#additional_encryption_context RdsIntegration#additional_encryption_context}
	AdditionalEncryptionContext *map[string]*string `field:"optional" json:"additionalEncryptionContext" yaml:"additionalEncryptionContext"`
	// Data filters for the integration.
	//
	// These filters determine which tables from the source database are sent to the target Amazon Redshift data warehouse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_integration#data_filter RdsIntegration#data_filter}
	DataFilter *string `field:"optional" json:"dataFilter" yaml:"dataFilter"`
	// A description of the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_integration#description RdsIntegration#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_integration#integration_name RdsIntegration#integration_name}
	IntegrationName *string `field:"optional" json:"integrationName" yaml:"integrationName"`
	// The AWS Key Management System (AWS KMS) key identifier for the key to use to encrypt the integration.
	//
	// If you don't specify an encryption key, RDS uses a default AWS owned key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_integration#kms_key_id RdsIntegration#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// A list of tags. For more information, see [Tagging Amazon RDS Resources](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in the *Amazon RDS User Guide.*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_integration#tags RdsIntegration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

