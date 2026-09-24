// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentagentspace


type SecurityagentAgentSpaceAwsResources struct {
	// IAM role ARNs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityagent_agent_space#iam_roles SecurityagentAgentSpace#iam_roles}
	IamRoles *[]*string `field:"optional" json:"iamRoles" yaml:"iamRoles"`
	// Lambda function ARNs used to retrieve tester credentials for pentests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityagent_agent_space#lambda_function_arns SecurityagentAgentSpace#lambda_function_arns}
	LambdaFunctionArns *[]*string `field:"optional" json:"lambdaFunctionArns" yaml:"lambdaFunctionArns"`
	// CloudWatch log group ARNs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityagent_agent_space#log_groups SecurityagentAgentSpace#log_groups}
	LogGroups *[]*string `field:"optional" json:"logGroups" yaml:"logGroups"`
	// S3 bucket ARNs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityagent_agent_space#s3_buckets SecurityagentAgentSpace#s3_buckets}
	S3Buckets *[]*string `field:"optional" json:"s3Buckets" yaml:"s3Buckets"`
	// SecretsManager secret ARNs used to store tester credentials for pentests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityagent_agent_space#secret_arns SecurityagentAgentSpace#secret_arns}
	SecretArns *[]*string `field:"optional" json:"secretArns" yaml:"secretArns"`
	// VPC configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityagent_agent_space#vpcs SecurityagentAgentSpace#vpcs}
	Vpcs interface{} `field:"optional" json:"vpcs" yaml:"vpcs"`
}

