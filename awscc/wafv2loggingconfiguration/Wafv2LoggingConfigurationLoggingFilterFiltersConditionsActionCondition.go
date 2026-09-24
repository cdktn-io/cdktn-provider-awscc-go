// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2loggingconfiguration


type Wafv2LoggingConfigurationLoggingFilterFiltersConditionsActionCondition struct {
	// Logic to apply to the filtering conditions.
	//
	// You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wafv2_logging_configuration#action Wafv2LoggingConfiguration#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
}

