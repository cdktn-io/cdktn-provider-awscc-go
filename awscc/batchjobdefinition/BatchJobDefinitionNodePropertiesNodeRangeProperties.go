// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionNodePropertiesNodeRangeProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/batch_job_definition#consumable_resource_properties BatchJobDefinition#consumable_resource_properties}.
	ConsumableResourceProperties *BatchJobDefinitionNodePropertiesNodeRangePropertiesConsumableResourceProperties `field:"optional" json:"consumableResourceProperties" yaml:"consumableResourceProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/batch_job_definition#container BatchJobDefinition#container}.
	Container *BatchJobDefinitionNodePropertiesNodeRangePropertiesContainer `field:"optional" json:"container" yaml:"container"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/batch_job_definition#ecs_properties BatchJobDefinition#ecs_properties}.
	EcsProperties *BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsProperties `field:"optional" json:"ecsProperties" yaml:"ecsProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/batch_job_definition#eks_properties BatchJobDefinition#eks_properties}.
	EksProperties *BatchJobDefinitionNodePropertiesNodeRangePropertiesEksProperties `field:"optional" json:"eksProperties" yaml:"eksProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/batch_job_definition#instance_types BatchJobDefinition#instance_types}.
	InstanceTypes *[]*string `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/batch_job_definition#target_nodes BatchJobDefinition#target_nodes}.
	TargetNodes *string `field:"optional" json:"targetNodes" yaml:"targetNodes"`
}

