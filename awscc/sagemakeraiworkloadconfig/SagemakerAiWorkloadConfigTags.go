// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeraiworkloadconfig


type SagemakerAiWorkloadConfigTags struct {
	// The tag key. Tag keys must be unique per resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_ai_workload_config#key SagemakerAiWorkloadConfig#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_ai_workload_config#value SagemakerAiWorkloadConfig#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

