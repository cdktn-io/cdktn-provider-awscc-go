// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskset


type EcsTaskSetServiceRegistries struct {
	// The container name value, already specified in the task definition, to be used for your service discovery service.
	//
	// If the task definition that your service task specifies uses the bridge or host network mode, you must specify a containerName and containerPort combination from the task definition. If the task definition that your service task specifies uses the awsvpc network mode and a type SRV DNS record is used, you must specify either a containerName and containerPort combination or a port value, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_task_set#container_name EcsTaskSet#container_name}
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The port value, already specified in the task definition, to be used for your service discovery service.
	//
	// If the task definition your service task specifies uses the bridge or host network mode, you must specify a containerName and containerPort combination from the task definition. If the task definition your service task specifies uses the awsvpc network mode and a type SRV DNS record is used, you must specify either a containerName and containerPort combination or a port value, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_task_set#container_port EcsTaskSet#container_port}
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// The port value used if your service discovery service specified an SRV record.
	//
	// This field may be used if both the awsvpc network mode and SRV records are used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_task_set#port EcsTaskSet#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The Amazon Resource Name (ARN) of the service registry.
	//
	// The currently supported service registry is AWS Cloud Map. For more information, see https://docs.aws.amazon.com/cloud-map/latest/api/API_CreateService.html
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_task_set#registry_arn EcsTaskSet#registry_arn}
	RegistryArn *string `field:"optional" json:"registryArn" yaml:"registryArn"`
}

