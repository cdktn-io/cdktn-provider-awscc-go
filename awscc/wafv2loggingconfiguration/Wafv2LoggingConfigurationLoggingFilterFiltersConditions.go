// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2loggingconfiguration


type Wafv2LoggingConfigurationLoggingFilterFiltersConditions struct {
	// A single action condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wafv2_logging_configuration#action_condition Wafv2LoggingConfiguration#action_condition}
	ActionCondition *Wafv2LoggingConfigurationLoggingFilterFiltersConditionsActionCondition `field:"optional" json:"actionCondition" yaml:"actionCondition"`
	// A single label name condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wafv2_logging_configuration#label_name_condition Wafv2LoggingConfiguration#label_name_condition}
	LabelNameCondition *Wafv2LoggingConfigurationLoggingFilterFiltersConditionsLabelNameCondition `field:"optional" json:"labelNameCondition" yaml:"labelNameCondition"`
}

