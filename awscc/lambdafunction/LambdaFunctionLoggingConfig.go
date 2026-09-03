// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionLoggingConfig struct {
	// Set this property to filter the application logs for your function that Lambda sends to CloudWatch.
	//
	// Lambda only sends application logs at the selected level of detail and lower, where ``TRACE`` is the highest level and ``FATAL`` is the lowest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lambda_function#application_log_level LambdaFunction#application_log_level}
	ApplicationLogLevel *string `field:"optional" json:"applicationLogLevel" yaml:"applicationLogLevel"`
	// The format in which Lambda sends your function's application and system logs to CloudWatch.
	//
	// Select between plain text and structured JSON.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lambda_function#log_format LambdaFunction#log_format}
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
	// The name of the Amazon CloudWatch log group the function sends logs to.
	//
	// By default, Lambda functions send logs to a default log group named ``/aws/lambda/<function name>``. To use a different log group, enter an existing log group or enter a new log group name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lambda_function#log_group LambdaFunction#log_group}
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Set this property to filter the system logs for your function that Lambda sends to CloudWatch.
	//
	// Lambda only sends system logs at the selected level of detail and lower, where ``DEBUG`` is the highest level and ``WARN`` is the lowest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lambda_function#system_log_level LambdaFunction#system_log_level}
	SystemLogLevel *string `field:"optional" json:"systemLogLevel" yaml:"systemLogLevel"`
}

