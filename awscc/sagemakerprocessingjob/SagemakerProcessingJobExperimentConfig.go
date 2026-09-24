// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerprocessingjob


type SagemakerProcessingJobExperimentConfig struct {
	// The name of an existing experiment to associate with the trial component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_processing_job#experiment_name SagemakerProcessingJob#experiment_name}
	ExperimentName *string `field:"optional" json:"experimentName" yaml:"experimentName"`
	// The name of the experiment run to associate with the trial component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_processing_job#run_name SagemakerProcessingJob#run_name}
	RunName *string `field:"optional" json:"runName" yaml:"runName"`
	// The display name for the trial component.
	//
	// If this key isn't specified, the display name is the trial component name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_processing_job#trial_component_display_name SagemakerProcessingJob#trial_component_display_name}
	TrialComponentDisplayName *string `field:"optional" json:"trialComponentDisplayName" yaml:"trialComponentDisplayName"`
	// The name of an existing trial to associate the trial component with.
	//
	// If not specified, a new trial is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_processing_job#trial_name SagemakerProcessingJob#trial_name}
	TrialName *string `field:"optional" json:"trialName" yaml:"trialName"`
}

