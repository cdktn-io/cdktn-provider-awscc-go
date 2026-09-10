// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionPlacementConstraints struct {
	// A cluster query language expression to apply to the constraint.
	//
	// For more information, see [Cluster query language](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html) in the *Amazon Elastic Container Service Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_task_definition#expression EcsTaskDefinition#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// The type of constraint. The ``MemberOf`` constraint restricts selection to be from a group of valid candidates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_task_definition#type EcsTaskDefinition#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

