// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftgameservergroup


type GameliftGameServerGroupLaunchTemplate struct {
	// A unique identifier for an existing EC2 launch template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_game_server_group#launch_template_id GameliftGameServerGroup#launch_template_id}
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// A readable identifier for an existing EC2 launch template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_game_server_group#launch_template_name GameliftGameServerGroup#launch_template_name}
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// The version of the EC2 launch template to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_game_server_group#version GameliftGameServerGroup#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

