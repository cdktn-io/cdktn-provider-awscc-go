// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferenceexperiment


type SagemakerInferenceExperimentSchedule struct {
	// The timestamp at which the inference experiment ended or will end.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_inference_experiment#end_time SagemakerInferenceExperiment#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The timestamp at which the inference experiment started or will start.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_inference_experiment#start_time SagemakerInferenceExperiment#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

