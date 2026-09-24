// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ec2fleet


type Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ec2_fleet#configuration_type Ec2Ec2Fleet#configuration_type}.
	ConfigurationType *string `field:"optional" json:"configurationType" yaml:"configurationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ec2_fleet#standby_availability_zones Ec2Ec2Fleet#standby_availability_zones}.
	StandbyAvailabilityZones interface{} `field:"optional" json:"standbyAvailabilityZones" yaml:"standbyAvailabilityZones"`
}

