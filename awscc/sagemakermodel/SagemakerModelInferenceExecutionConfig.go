// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodel


type SagemakerModelInferenceExecutionConfig struct {
	// How containers in a multi-container are run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_model#mode SagemakerModel#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

