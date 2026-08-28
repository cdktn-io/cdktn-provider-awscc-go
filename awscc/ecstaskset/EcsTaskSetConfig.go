// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsTaskSetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_set#cluster EcsTaskSet#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The short name or full Amazon Resource Name (ARN) of the service to create the task set in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_set#service EcsTaskSet#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// The short name or full Amazon Resource Name (ARN) of the task definition for the tasks in the task set to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_set#task_definition EcsTaskSet#task_definition}
	TaskDefinition *string `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_set#capacity_provider_strategy EcsTaskSet#capacity_provider_strategy}.
	CapacityProviderStrategy interface{} `field:"optional" json:"capacityProviderStrategy" yaml:"capacityProviderStrategy"`
	// An optional non-unique tag that identifies this task set in external systems.
	//
	// If the task set is associated with a service discovery registry, the tasks in this task set will have the ECS_TASK_SET_EXTERNAL_ID AWS Cloud Map attribute set to the provided value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_set#external_id EcsTaskSet#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// The launch type that new tasks in the task set will use.
	//
	// For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html in the Amazon Elastic Container Service Developer Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_set#launch_type EcsTaskSet#launch_type}
	LaunchType *string `field:"optional" json:"launchType" yaml:"launchType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_set#load_balancers EcsTaskSet#load_balancers}.
	LoadBalancers interface{} `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// An object representing the network configuration for a task or service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_set#network_configuration EcsTaskSet#network_configuration}
	NetworkConfiguration *EcsTaskSetNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The platform version that the tasks in the task set should use.
	//
	// A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the LATEST platform version is used by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_set#platform_version EcsTaskSet#platform_version}
	PlatformVersion *string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// A floating-point percentage of the desired number of tasks to place and keep running in the task set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_set#scale EcsTaskSet#scale}
	Scale *EcsTaskSetScale `field:"optional" json:"scale" yaml:"scale"`
	// The details of the service discovery registries to assign to this task set. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_set#service_registries EcsTaskSet#service_registries}
	ServiceRegistries interface{} `field:"optional" json:"serviceRegistries" yaml:"serviceRegistries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_set#tags EcsTaskSet#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

