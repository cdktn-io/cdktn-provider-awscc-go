// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerexperimenttrialcomponent


type SagemakerExperimentTrialComponentStatus struct {
	// If the component failed, a message describing why.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_experiment_trial_component#message SagemakerExperimentTrialComponent#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
	// The status of the trial component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_experiment_trial_component#primary_status SagemakerExperimentTrialComponent#primary_status}
	PrimaryStatus *string `field:"optional" json:"primaryStatus" yaml:"primaryStatus"`
}

