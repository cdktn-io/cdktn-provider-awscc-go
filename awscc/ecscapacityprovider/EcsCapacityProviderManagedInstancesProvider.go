// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecscapacityprovider


type EcsCapacityProviderManagedInstancesProvider struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_capacity_provider#auto_repair_configuration EcsCapacityProvider#auto_repair_configuration}.
	AutoRepairConfiguration *EcsCapacityProviderManagedInstancesProviderAutoRepairConfiguration `field:"optional" json:"autoRepairConfiguration" yaml:"autoRepairConfiguration"`
	// Defines how Amazon ECS Managed Instances optimizes the infrastructure in your capacity provider.
	//
	// Configure it to turn on or off the infrastructure optimization in your capacity provider, and to control the idle EC2 instances optimization delay.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_capacity_provider#infrastructure_optimization EcsCapacityProvider#infrastructure_optimization}
	InfrastructureOptimization *EcsCapacityProviderManagedInstancesProviderInfrastructureOptimization `field:"optional" json:"infrastructureOptimization" yaml:"infrastructureOptimization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_capacity_provider#infrastructure_role_arn EcsCapacityProvider#infrastructure_role_arn}.
	InfrastructureRoleArn *string `field:"optional" json:"infrastructureRoleArn" yaml:"infrastructureRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_capacity_provider#instance_launch_template EcsCapacityProvider#instance_launch_template}.
	InstanceLaunchTemplate *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplate `field:"optional" json:"instanceLaunchTemplate" yaml:"instanceLaunchTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_capacity_provider#propagate_tags EcsCapacityProvider#propagate_tags}.
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
}

