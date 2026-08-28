// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2loggingconfiguration


type Wafv2LoggingConfigurationLoggingFilterFilters struct {
	// How to handle logs that satisfy the filter's conditions and requirement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wafv2_logging_configuration#behavior Wafv2LoggingConfiguration#behavior}
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
	// Match conditions for the filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wafv2_logging_configuration#conditions Wafv2LoggingConfiguration#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Logic to apply to the filtering conditions.
	//
	// You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wafv2_logging_configuration#requirement Wafv2LoggingConfiguration#requirement}
	Requirement *string `field:"optional" json:"requirement" yaml:"requirement"`
}

