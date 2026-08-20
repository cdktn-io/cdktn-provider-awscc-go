// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2spotfleet


type Ec2SpotFleetSpotFleetRequestConfigDataLoadBalancersConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_spot_fleet#classic_load_balancers_config Ec2SpotFleet#classic_load_balancers_config}.
	ClassicLoadBalancersConfig *Ec2SpotFleetSpotFleetRequestConfigDataLoadBalancersConfigClassicLoadBalancersConfig `field:"optional" json:"classicLoadBalancersConfig" yaml:"classicLoadBalancersConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_spot_fleet#target_groups_config Ec2SpotFleet#target_groups_config}.
	TargetGroupsConfig *Ec2SpotFleetSpotFleetRequestConfigDataLoadBalancersConfigTargetGroupsConfig `field:"optional" json:"targetGroupsConfig" yaml:"targetGroupsConfig"`
}

