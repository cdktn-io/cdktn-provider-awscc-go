// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceForceNewDeployment struct {
	// Determines whether to force a new deployment of the service.
	//
	// By default, deployments aren't forced. You can use this option to start a new deployment with no service definition changes. For example, you can update a service's tasks to use a newer Docker image with the same image/tag combination (``my_image:latest``) or to roll Fargate tasks onto a newer platform version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ecs_service#enable_force_new_deployment EcsService#enable_force_new_deployment}
	EnableForceNewDeployment interface{} `field:"optional" json:"enableForceNewDeployment" yaml:"enableForceNewDeployment"`
	// When you change the``ForceNewDeploymentNonce`` value in your template, it signals ECS to start a new deployment even though no other service parameters have changed.
	//
	// The value must be a unique, time- varying value like a timestamp, random string, or sequence number. Use this property when you want to ensure your tasks pick up the latest version of a Docker image that uses the same tag but has been updated in the registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ecs_service#force_new_deployment_nonce EcsService#force_new_deployment_nonce}
	ForceNewDeploymentNonce *string `field:"optional" json:"forceNewDeploymentNonce" yaml:"forceNewDeploymentNonce"`
}

