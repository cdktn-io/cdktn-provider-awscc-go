// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationresourceversion


type CloudformationResourceVersionLoggingConfig struct {
	// The Amazon CloudWatch log group to which CloudFormation sends error logging information when invoking the type's handlers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudformation_resource_version#log_group_name CloudformationResourceVersion#log_group_name}
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The ARN of the role that CloudFormation should assume when sending log entries to CloudWatch logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudformation_resource_version#log_role_arn CloudformationResourceVersion#log_role_arn}
	LogRoleArn *string `field:"optional" json:"logRoleArn" yaml:"logRoleArn"`
}

