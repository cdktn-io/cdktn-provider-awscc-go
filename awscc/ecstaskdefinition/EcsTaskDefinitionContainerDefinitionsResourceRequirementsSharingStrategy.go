// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsResourceRequirementsSharingStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_task_definition#nvidia_mps EcsTaskDefinition#nvidia_mps}.
	NvidiaMps *EcsTaskDefinitionContainerDefinitionsResourceRequirementsSharingStrategyNvidiaMps `field:"optional" json:"nvidiaMps" yaml:"nvidiaMps"`
}

