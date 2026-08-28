// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ec2fleet


type Ec2Ec2FleetLaunchTemplateConfigsLaunchTemplateSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_ec2_fleet#launch_template_id Ec2Ec2Fleet#launch_template_id}.
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_ec2_fleet#launch_template_name Ec2Ec2Fleet#launch_template_name}.
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_ec2_fleet#launch_template_specification_user_data Ec2Ec2Fleet#launch_template_specification_user_data}.
	LaunchTemplateSpecificationUserData *string `field:"optional" json:"launchTemplateSpecificationUserData" yaml:"launchTemplateSpecificationUserData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_ec2_fleet#version Ec2Ec2Fleet#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

