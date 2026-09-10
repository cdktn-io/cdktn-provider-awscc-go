// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrialcomponent


type SagemakerTrialComponentParameters struct {
	// The numeric value of a numeric hyperparameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_trial_component#number_value SagemakerTrialComponent#number_value}
	NumberValue *float64 `field:"optional" json:"numberValue" yaml:"numberValue"`
	// The string value of a categorical hyperparameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_trial_component#string_value SagemakerTrialComponent#string_value}
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

