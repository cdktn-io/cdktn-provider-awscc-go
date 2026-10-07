// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdawebfunctionrevision


type LambdaWebFunctionRevisionServiceConfigTelemetryConfig struct {
	// The logging configuration for the web function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/lambda_web_function_revision#logging_config LambdaWebFunctionRevision#logging_config}
	LoggingConfig *LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
}

