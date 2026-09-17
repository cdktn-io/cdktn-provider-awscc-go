// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftgameservergroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GameliftGameServerGroupConfig struct {
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
	// An identifier for the new game server group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_game_server_group#game_server_group_name GameliftGameServerGroup#game_server_group_name}
	GameServerGroupName *string `field:"required" json:"gameServerGroupName" yaml:"gameServerGroupName"`
	// A set of EC2 instance types to use when creating instances in the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_game_server_group#instance_definitions GameliftGameServerGroup#instance_definitions}
	InstanceDefinitions interface{} `field:"required" json:"instanceDefinitions" yaml:"instanceDefinitions"`
	// The Amazon Resource Name (ARN) for an IAM role that allows Amazon GameLift to access your EC2 Auto Scaling groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_game_server_group#role_arn GameliftGameServerGroup#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Configuration settings to define a scaling policy for the Auto Scaling group that is optimized for game hosting.
	//
	// Updating this game server group property will not take effect for the created EC2 Auto Scaling group, please update the EC2 Auto Scaling group directly after creating the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_game_server_group#auto_scaling_policy GameliftGameServerGroup#auto_scaling_policy}
	AutoScalingPolicy *GameliftGameServerGroupAutoScalingPolicy `field:"optional" json:"autoScalingPolicy" yaml:"autoScalingPolicy"`
	// The fallback balancing method to use for the game server group when Spot Instances in a Region become unavailable or are not viable for game hosting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_game_server_group#balancing_strategy GameliftGameServerGroup#balancing_strategy}
	BalancingStrategy *string `field:"optional" json:"balancingStrategy" yaml:"balancingStrategy"`
	// The type of delete to perform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_game_server_group#delete_option GameliftGameServerGroup#delete_option}
	DeleteOption *string `field:"optional" json:"deleteOption" yaml:"deleteOption"`
	// A flag that indicates whether instances in the game server group are protected from early termination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_game_server_group#game_server_protection_policy GameliftGameServerGroup#game_server_protection_policy}
	GameServerProtectionPolicy *string `field:"optional" json:"gameServerProtectionPolicy" yaml:"gameServerProtectionPolicy"`
	// The EC2 launch template that contains configuration settings and game server code to be deployed to all instances in the game server group.
	//
	// Updating this game server group property will not take effect for the created EC2 Auto Scaling group, please update the EC2 Auto Scaling group directly after creating the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_game_server_group#launch_template GameliftGameServerGroup#launch_template}
	LaunchTemplate *GameliftGameServerGroupLaunchTemplate `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// The maximum number of instances allowed in the EC2 Auto Scaling group.
	//
	// Updating this game server group property will not take effect for the created EC2 Auto Scaling group, please update the EC2 Auto Scaling group directly after creating the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_game_server_group#max_size GameliftGameServerGroup#max_size}
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The minimum number of instances allowed in the EC2 Auto Scaling group.
	//
	// Updating this game server group property will not take effect for the created EC2 Auto Scaling group, please update the EC2 Auto Scaling group directly after creating the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_game_server_group#min_size GameliftGameServerGroup#min_size}
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// A list of labels to assign to the new game server group resource.
	//
	// Updating game server group tags with CloudFormation will not take effect. Please update this property using AWS GameLift APIs instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_game_server_group#tags GameliftGameServerGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A list of virtual private cloud (VPC) subnets to use with instances in the game server group.
	//
	// Updating this game server group property will not take effect for the created EC2 Auto Scaling group, please update the EC2 Auto Scaling group directly after creating the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_game_server_group#vpc_subnets GameliftGameServerGroup#vpc_subnets}
	VpcSubnets *[]*string `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

