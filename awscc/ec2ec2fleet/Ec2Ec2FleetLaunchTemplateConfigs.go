// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ec2fleet


type Ec2Ec2FleetLaunchTemplateConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_ec2_fleet#launch_template_specification Ec2Ec2Fleet#launch_template_specification}.
	LaunchTemplateSpecification *Ec2Ec2FleetLaunchTemplateConfigsLaunchTemplateSpecification `field:"optional" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_ec2_fleet#overrides Ec2Ec2Fleet#overrides}.
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
}

