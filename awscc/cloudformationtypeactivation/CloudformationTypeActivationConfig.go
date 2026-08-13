// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationtypeactivation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudformationTypeActivationConfig struct {
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
	// Whether to automatically update the extension in this account and region when a new minor version is published by the extension publisher.
	//
	// Major versions released by the publisher must be manually updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_type_activation#auto_update CloudformationTypeActivation#auto_update}
	AutoUpdate interface{} `field:"optional" json:"autoUpdate" yaml:"autoUpdate"`
	// The Amazon Resource Name (ARN) of the IAM execution role to use to register the type.
	//
	// If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_type_activation#execution_role_arn CloudformationTypeActivation#execution_role_arn}
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Specifies logging configuration information for a type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_type_activation#logging_config CloudformationTypeActivation#logging_config}
	LoggingConfig *CloudformationTypeActivationLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// The Major Version of the type you want to enable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_type_activation#major_version CloudformationTypeActivation#major_version}
	MajorVersion *string `field:"optional" json:"majorVersion" yaml:"majorVersion"`
	// The Amazon Resource Number (ARN) assigned to the public extension upon publication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_type_activation#public_type_arn CloudformationTypeActivation#public_type_arn}
	PublicTypeArn *string `field:"optional" json:"publicTypeArn" yaml:"publicTypeArn"`
	// The reserved publisher id for this type, or the publisher id assigned by CloudFormation for publishing in this region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_type_activation#publisher_id CloudformationTypeActivation#publisher_id}
	PublisherId *string `field:"optional" json:"publisherId" yaml:"publisherId"`
	// The kind of extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_type_activation#type CloudformationTypeActivation#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_type_activation#type_name CloudformationTypeActivation#type_name}
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
	// An alias to assign to the public extension in this account and region.
	//
	// If you specify an alias for the extension, you must then use the alias to refer to the extension in your templates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_type_activation#type_name_alias CloudformationTypeActivation#type_name_alias}
	TypeNameAlias *string `field:"optional" json:"typeNameAlias" yaml:"typeNameAlias"`
	// Manually updates a previously-enabled type to a new major or minor version, if available.
	//
	// You can also use this parameter to update the value of AutoUpdateEnabled
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudformation_type_activation#version_bump CloudformationTypeActivation#version_bump}
	VersionBump *string `field:"optional" json:"versionBump" yaml:"versionBump"`
}

