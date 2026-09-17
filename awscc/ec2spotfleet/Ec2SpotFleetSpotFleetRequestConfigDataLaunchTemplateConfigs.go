// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2spotfleet


type Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_spot_fleet#launch_template_specification Ec2SpotFleet#launch_template_specification}.
	LaunchTemplateSpecification *Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsLaunchTemplateSpecification `field:"optional" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_spot_fleet#overrides Ec2SpotFleet#overrides}.
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
}

