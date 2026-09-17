// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesconfigurationset


type SesConfigurationSetSuppressionOptions struct {
	// A list that contains the reasons that email addresses are automatically added to the suppression list for your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_configuration_set#suppressed_reasons SesConfigurationSet#suppressed_reasons}
	SuppressedReasons *[]*string `field:"optional" json:"suppressedReasons" yaml:"suppressedReasons"`
	// An object that contains information about the validation options for your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_configuration_set#validation_options SesConfigurationSet#validation_options}
	ValidationOptions *SesConfigurationSetSuppressionOptionsValidationOptions `field:"optional" json:"validationOptions" yaml:"validationOptions"`
}

