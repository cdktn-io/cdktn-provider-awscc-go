// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemon

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsDaemonConfig struct {
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
	// The Amazon Resource Names (ARNs) of the capacity providers associated with the daemon.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_daemon#capacity_provider_arns EcsDaemon#capacity_provider_arns}
	CapacityProviderArns *[]*string `field:"optional" json:"capacityProviderArns" yaml:"capacityProviderArns"`
	// The Amazon Resource Name (ARN) of the cluster that the daemon is running in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_daemon#cluster_arn EcsDaemon#cluster_arn}
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_daemon#daemon_name EcsDaemon#daemon_name}.
	DaemonName *string `field:"optional" json:"daemonName" yaml:"daemonName"`
	// The Amazon Resource Name (ARN) of the daemon task definition used by this revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_daemon#daemon_task_definition_arn EcsDaemon#daemon_task_definition_arn}
	DaemonTaskDefinitionArn *string `field:"optional" json:"daemonTaskDefinitionArn" yaml:"daemonTaskDefinitionArn"`
	// The deployment configuration used for this daemon deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_daemon#deployment_configuration EcsDaemon#deployment_configuration}
	DeploymentConfiguration *EcsDaemonDeploymentConfiguration `field:"optional" json:"deploymentConfiguration" yaml:"deploymentConfiguration"`
	// Specifies whether Amazon ECS managed tags are turned on for the daemon tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_daemon#enable_ecs_managed_tags EcsDaemon#enable_ecs_managed_tags}
	EnableEcsManagedTags interface{} `field:"optional" json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// Specifies whether the execute command functionality is turned on for the daemon tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_daemon#enable_execute_command EcsDaemon#enable_execute_command}
	EnableExecuteCommand interface{} `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// Specifies whether tags are propagated from the daemon to the daemon tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_daemon#propagate_tags EcsDaemon#propagate_tags}
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_daemon#tags EcsDaemon#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

