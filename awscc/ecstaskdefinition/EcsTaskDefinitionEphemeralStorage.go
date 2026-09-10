// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionEphemeralStorage struct {
	// The total amount, in GiB, of ephemeral storage to set for the task.
	//
	// The minimum supported value is ``21`` GiB and the maximum supported value is ``200`` GiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_task_definition#size_in_gi_b EcsTaskDefinition#size_in_gi_b}
	SizeInGiB *float64 `field:"optional" json:"sizeInGiB" yaml:"sizeInGiB"`
}

