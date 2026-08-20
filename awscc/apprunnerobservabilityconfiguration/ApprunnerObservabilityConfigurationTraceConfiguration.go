// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apprunnerobservabilityconfiguration


type ApprunnerObservabilityConfigurationTraceConfiguration struct {
	// The implementation provider chosen for tracing App Runner services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apprunner_observability_configuration#vendor ApprunnerObservabilityConfiguration#vendor}
	Vendor *string `field:"optional" json:"vendor" yaml:"vendor"`
}

