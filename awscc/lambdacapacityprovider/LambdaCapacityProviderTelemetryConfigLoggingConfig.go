// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdacapacityprovider


type LambdaCapacityProviderTelemetryConfigLoggingConfig struct {
	// The name of the Amazon CloudWatch log group the capacity provider sends logs to.
	//
	// By default, Lambda capacity providers send logs to a default log group named ``/aws/lambda/capacity-provider/<capacity provider name>``. To use a different log group, enter an existing log group or enter a new log group name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_capacity_provider#log_group LambdaCapacityProvider#log_group}
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Set this property to filter the system logs for your capacity provider that Lambda sends to CloudWatch.
	//
	// Lambda only sends system logs at the selected level of detail and lower, where ``DEBUG`` is the highest level and ``WARN`` is the lowest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_capacity_provider#system_log_level LambdaCapacityProvider#system_log_level}
	SystemLogLevel *string `field:"optional" json:"systemLogLevel" yaml:"systemLogLevel"`
}

