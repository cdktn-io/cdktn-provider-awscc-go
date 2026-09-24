// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrialcomponent


type SagemakerTrialComponentStatus struct {
	// If the component failed, a message describing why.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_trial_component#message SagemakerTrialComponent#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
	// The status of the trial component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_trial_component#primary_status SagemakerTrialComponent#primary_status}
	PrimaryStatus *string `field:"optional" json:"primaryStatus" yaml:"primaryStatus"`
}

