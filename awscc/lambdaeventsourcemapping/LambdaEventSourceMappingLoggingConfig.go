// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdaeventsourcemapping


type LambdaEventSourceMappingLoggingConfig struct {
	// Set this property to filter the system logs for your function that Lambda sends to CloudWatch.
	//
	// Lambda only sends system logs at the selected level of detail and lower, where ``DEBUG`` is the highest level and ``WARN`` is the lowest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_event_source_mapping#system_log_level LambdaEventSourceMapping#system_log_level}
	SystemLogLevel *string `field:"optional" json:"systemLogLevel" yaml:"systemLogLevel"`
}

