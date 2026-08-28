// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesconfigurationset


type SesConfigurationSetSuppressionOptionsValidationOptions struct {
	// The condition threshold settings for suppression validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_configuration_set#condition_threshold SesConfigurationSet#condition_threshold}
	ConditionThreshold *SesConfigurationSetSuppressionOptionsValidationOptionsConditionThreshold `field:"optional" json:"conditionThreshold" yaml:"conditionThreshold"`
}

