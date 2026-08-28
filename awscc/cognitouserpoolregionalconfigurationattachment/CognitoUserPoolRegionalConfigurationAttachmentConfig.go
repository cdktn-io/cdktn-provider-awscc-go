// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitouserpoolregionalconfigurationattachment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CognitoUserPoolRegionalConfigurationAttachmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_user_pool_regional_configuration_attachment#user_pool_id CognitoUserPoolRegionalConfigurationAttachment#user_pool_id}.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_user_pool_regional_configuration_attachment#email_configuration CognitoUserPoolRegionalConfigurationAttachment#email_configuration}.
	EmailConfiguration *CognitoUserPoolRegionalConfigurationAttachmentEmailConfiguration `field:"optional" json:"emailConfiguration" yaml:"emailConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_user_pool_regional_configuration_attachment#lambda_config CognitoUserPoolRegionalConfigurationAttachment#lambda_config}.
	LambdaConfig *CognitoUserPoolRegionalConfigurationAttachmentLambdaConfig `field:"optional" json:"lambdaConfig" yaml:"lambdaConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_user_pool_regional_configuration_attachment#sms_configuration CognitoUserPoolRegionalConfigurationAttachment#sms_configuration}.
	SmsConfiguration *CognitoUserPoolRegionalConfigurationAttachmentSmsConfiguration `field:"optional" json:"smsConfiguration" yaml:"smsConfiguration"`
	// The status of the replica. Set to ACTIVE or INACTIVE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_user_pool_regional_configuration_attachment#status CognitoUserPoolRegionalConfigurationAttachment#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_user_pool_regional_configuration_attachment#user_pool_tags CognitoUserPoolRegionalConfigurationAttachment#user_pool_tags}.
	UserPoolTags *map[string]*string `field:"optional" json:"userPoolTags" yaml:"userPoolTags"`
}

