// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdacapacityprovider


type LambdaCapacityProviderTelemetryConfig struct {
	// The capacity provider's Amazon CloudWatch Logs configuration settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_capacity_provider#logging_config LambdaCapacityProvider#logging_config}
	LoggingConfig *LambdaCapacityProviderTelemetryConfigLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
}

