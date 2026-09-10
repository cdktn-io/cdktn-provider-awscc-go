// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrialcomponent


type SagemakerTrialComponentInputArtifacts struct {
	// The media type of the artifact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_trial_component#media_type SagemakerTrialComponent#media_type}
	MediaType *string `field:"optional" json:"mediaType" yaml:"mediaType"`
	// The location of the artifact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_trial_component#value SagemakerTrialComponent#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

