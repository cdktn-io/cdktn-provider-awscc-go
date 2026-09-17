// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelqualityjobdefinition


type SagemakerModelQualityJobDefinitionStoppingCondition struct {
	// The maximum runtime allowed in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_quality_job_definition#max_runtime_in_seconds SagemakerModelQualityJobDefinition#max_runtime_in_seconds}
	MaxRuntimeInSeconds *float64 `field:"optional" json:"maxRuntimeInSeconds" yaml:"maxRuntimeInSeconds"`
}

