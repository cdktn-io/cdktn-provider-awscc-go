// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsResourceRequirementsSharingStrategyNvidiaMps struct {
	// The maximum percentage of the GPU's compute capacity that the container can use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_task_definition#max_compute_percent EcsTaskDefinition#max_compute_percent}
	MaxComputePercent *float64 `field:"optional" json:"maxComputePercent" yaml:"maxComputePercent"`
	// The amount of GPU memory, in MiB, allocated to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_task_definition#memory EcsTaskDefinition#memory}
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
}

