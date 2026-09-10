// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationresourceversion

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudformationResourceVersionConfig struct {
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
	// A url to the S3 bucket containing the schema handler package that contains the schema, event handlers, and associated files for the type you want to register.
	//
	// For information on generating a schema handler package for the type you want to register, see submit in the CloudFormation CLI User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudformation_resource_version#schema_handler_package CloudformationResourceVersion#schema_handler_package}
	SchemaHandlerPackage *string `field:"required" json:"schemaHandlerPackage" yaml:"schemaHandlerPackage"`
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudformation_resource_version#type_name CloudformationResourceVersion#type_name}
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// The Amazon Resource Name (ARN) of the IAM execution role to use to register the type.
	//
	// If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudformation_resource_version#execution_role_arn CloudformationResourceVersion#execution_role_arn}
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Specifies logging configuration information for a type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudformation_resource_version#logging_config CloudformationResourceVersion#logging_config}
	LoggingConfig *CloudformationResourceVersionLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
}

