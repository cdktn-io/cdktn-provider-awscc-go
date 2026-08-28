// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationtypeactivation


type CloudformationTypeActivationLoggingConfig struct {
	// The Amazon CloudWatch log group to which CloudFormation sends error logging information when invoking the type's handlers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_type_activation#log_group_name CloudformationTypeActivation#log_group_name}
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The ARN of the role that CloudFormation should assume when sending log entries to CloudWatch logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_type_activation#log_role_arn CloudformationTypeActivation#log_role_arn}
	LogRoleArn *string `field:"optional" json:"logRoleArn" yaml:"logRoleArn"`
}

