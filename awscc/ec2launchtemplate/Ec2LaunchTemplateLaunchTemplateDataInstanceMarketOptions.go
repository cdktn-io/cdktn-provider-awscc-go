// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptions struct {
	// The market type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_launch_template#market_type Ec2LaunchTemplate#market_type}
	MarketType *string `field:"optional" json:"marketType" yaml:"marketType"`
	// The options for Spot Instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_launch_template#spot_options Ec2LaunchTemplate#spot_options}
	SpotOptions *Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptions `field:"optional" json:"spotOptions" yaml:"spotOptions"`
}

