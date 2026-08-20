// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainerfleet

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GameliftContainerFleetConfig struct {
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
	// A unique identifier for an AWS IAM role that manages access to your AWS services.
	//
	// Create a role or look up a role's ARN from the IAM dashboard in the AWS Management Console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#fleet_role_arn GameliftContainerFleet#fleet_role_arn}
	FleetRoleArn *string `field:"required" json:"fleetRoleArn" yaml:"fleetRoleArn"`
	// Indicates whether to use On-Demand instances or Spot instances for this fleet.
	//
	// If empty, the default is ON_DEMAND. Both categories of instances use identical hardware and configurations based on the instance type selected for this fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#billing_type GameliftContainerFleet#billing_type}
	BillingType *string `field:"optional" json:"billingType" yaml:"billingType"`
	// Provides details about how to drain old tasks and replace them with new updated tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#deployment_configuration GameliftContainerFleet#deployment_configuration}
	DeploymentConfiguration *GameliftContainerFleetDeploymentConfiguration `field:"optional" json:"deploymentConfiguration" yaml:"deploymentConfiguration"`
	// A human-readable description of a fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#description GameliftContainerFleet#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the container group definition that will be created per game server.
	//
	// You must specify GAME_SERVER container group. You have the option to also specify one PER_INSTANCE container group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#game_server_container_group_definition_name GameliftContainerFleet#game_server_container_group_definition_name}
	GameServerContainerGroupDefinitionName *string `field:"optional" json:"gameServerContainerGroupDefinitionName" yaml:"gameServerContainerGroupDefinitionName"`
	// The number of desired game server container groups per instance, a number between 1-5000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#game_server_container_groups_per_instance GameliftContainerFleet#game_server_container_groups_per_instance}
	GameServerContainerGroupsPerInstance *float64 `field:"optional" json:"gameServerContainerGroupsPerInstance" yaml:"gameServerContainerGroupsPerInstance"`
	// A policy that limits the number of game sessions an individual player can create over a span of time for this fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#game_session_creation_limit_policy GameliftContainerFleet#game_session_creation_limit_policy}
	GameSessionCreationLimitPolicy *GameliftContainerFleetGameSessionCreationLimitPolicy `field:"optional" json:"gameSessionCreationLimitPolicy" yaml:"gameSessionCreationLimitPolicy"`
	// Defines the range of ports on the instance that allow inbound traffic to connect with containers in a fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#instance_connection_port_range GameliftContainerFleet#instance_connection_port_range}
	InstanceConnectionPortRange *GameliftContainerFleetInstanceConnectionPortRange `field:"optional" json:"instanceConnectionPortRange" yaml:"instanceConnectionPortRange"`
	// A range of IP addresses and port settings that allow inbound traffic to connect to server processes on an Amazon GameLift server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#instance_inbound_permissions GameliftContainerFleet#instance_inbound_permissions}
	InstanceInboundPermissions interface{} `field:"optional" json:"instanceInboundPermissions" yaml:"instanceInboundPermissions"`
	// The name of an EC2 instance type that is supported in Amazon GameLift.
	//
	// A fleet instance type determines the computing resources of each instance in the fleet, including CPU, memory, storage, and networking capacity. Amazon GameLift supports the following EC2 instance types. See Amazon EC2 Instance Types for detailed descriptions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#instance_type GameliftContainerFleet#instance_type}
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#locations GameliftContainerFleet#locations}.
	Locations interface{} `field:"optional" json:"locations" yaml:"locations"`
	// A policy the location and provider of logs from the fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#log_configuration GameliftContainerFleet#log_configuration}
	LogConfiguration *GameliftContainerFleetLogConfiguration `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The name of an Amazon CloudWatch metric group.
	//
	// A metric group aggregates the metrics for all fleets in the group. Specify a string containing the metric group name. You can use an existing name or use a new name to create a new metric group. Currently, this parameter can have only one string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#metric_groups GameliftContainerFleet#metric_groups}
	MetricGroups *[]*string `field:"optional" json:"metricGroups" yaml:"metricGroups"`
	// A game session protection policy to apply to all game sessions hosted on instances in this fleet.
	//
	// When protected, active game sessions cannot be terminated during a scale-down event. If this parameter is not set, instances in this fleet default to no protection. You can change a fleet's protection policy to affect future game sessions on the fleet. You can also set protection for individual game sessions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#new_game_session_protection_policy GameliftContainerFleet#new_game_session_protection_policy}
	NewGameSessionProtectionPolicy *string `field:"optional" json:"newGameSessionProtectionPolicy" yaml:"newGameSessionProtectionPolicy"`
	// The name of the container group definition that will be created per instance.
	//
	// This field is optional if you specify GameServerContainerGroupDefinitionName.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#per_instance_container_group_definition_name GameliftContainerFleet#per_instance_container_group_definition_name}
	PerInstanceContainerGroupDefinitionName *string `field:"optional" json:"perInstanceContainerGroupDefinitionName" yaml:"perInstanceContainerGroupDefinitionName"`
	// The player gateway mode for the container fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#player_gateway_mode GameliftContainerFleet#player_gateway_mode}
	PlayerGatewayMode *string `field:"optional" json:"playerGatewayMode" yaml:"playerGatewayMode"`
	// A list of rules that control how a fleet is scaled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#scaling_policies GameliftContainerFleet#scaling_policies}
	ScalingPolicies interface{} `field:"optional" json:"scalingPolicies" yaml:"scalingPolicies"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_container_fleet#tags GameliftContainerFleet#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

