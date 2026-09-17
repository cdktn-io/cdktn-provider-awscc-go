// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesconfigurationset


type SesConfigurationSetSuppressionOptionsValidationOptionsConditionThreshold struct {
	// Whether the condition threshold is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_configuration_set#condition_threshold_enabled SesConfigurationSet#condition_threshold_enabled}
	ConditionThresholdEnabled *string `field:"optional" json:"conditionThresholdEnabled" yaml:"conditionThresholdEnabled"`
	// The overall confidence threshold settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_configuration_set#overall_confidence_threshold SesConfigurationSet#overall_confidence_threshold}
	OverallConfidenceThreshold *SesConfigurationSetSuppressionOptionsValidationOptionsConditionThresholdOverallConfidenceThreshold `field:"optional" json:"overallConfidenceThreshold" yaml:"overallConfidenceThreshold"`
}

