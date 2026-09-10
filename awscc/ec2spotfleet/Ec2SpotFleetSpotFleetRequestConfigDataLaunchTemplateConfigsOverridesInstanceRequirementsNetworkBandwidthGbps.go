// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2spotfleet


type Ec2SpotFleetSpotFleetRequestConfigDataLaunchTemplateConfigsOverridesInstanceRequirementsNetworkBandwidthGbps struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_spot_fleet#max Ec2SpotFleet#max}.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_spot_fleet#min Ec2SpotFleet#min}.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

