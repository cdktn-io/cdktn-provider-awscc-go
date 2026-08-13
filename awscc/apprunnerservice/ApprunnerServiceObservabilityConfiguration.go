// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apprunnerservice


type ApprunnerServiceObservabilityConfiguration struct {
	// The Amazon Resource Name (ARN) of the App Runner ObservabilityConfiguration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/apprunner_service#observability_configuration_arn ApprunnerService#observability_configuration_arn}
	ObservabilityConfigurationArn *string `field:"optional" json:"observabilityConfigurationArn" yaml:"observabilityConfigurationArn"`
	// Observability enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/apprunner_service#observability_enabled ApprunnerService#observability_enabled}
	ObservabilityEnabled interface{} `field:"optional" json:"observabilityEnabled" yaml:"observabilityEnabled"`
}

