// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesconfigurationset


type SesConfigurationSetSuppressionOptionsValidationOptionsConditionThresholdOverallConfidenceThreshold struct {
	// The confidence verdict threshold level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ses_configuration_set#confidence_verdict_threshold SesConfigurationSet#confidence_verdict_threshold}
	ConfidenceVerdictThreshold *string `field:"optional" json:"confidenceVerdictThreshold" yaml:"confidenceVerdictThreshold"`
}

