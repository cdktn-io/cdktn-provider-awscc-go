// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelbiasjobdefinition


type SagemakerModelBiasJobDefinitionStoppingCondition struct {
	// The maximum runtime allowed in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_bias_job_definition#max_runtime_in_seconds SagemakerModelBiasJobDefinition#max_runtime_in_seconds}
	MaxRuntimeInSeconds *float64 `field:"optional" json:"maxRuntimeInSeconds" yaml:"maxRuntimeInSeconds"`
}

