// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ec2fleet


type Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsMemoryMiB struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_ec2_fleet#max Ec2Ec2Fleet#max}.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_ec2_fleet#min Ec2Ec2Fleet#min}.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

