// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftfleet


type GameliftFleetLocationsLocationCapacity struct {
	// Defaults to MinSize if not defined.
	//
	// The number of EC2 instances you want to maintain in the specified fleet location. This value must fall between the minimum and maximum size limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/gamelift_fleet#desired_ec2_instances GameliftFleet#desired_ec2_instances}
	DesiredEc2Instances *float64 `field:"optional" json:"desiredEc2Instances" yaml:"desiredEc2Instances"`
	// Configuration options for Amazon GameLift Servers-managed capacity behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/gamelift_fleet#managed_capacity_configuration GameliftFleet#managed_capacity_configuration}
	ManagedCapacityConfiguration *GameliftFleetLocationsLocationCapacityManagedCapacityConfiguration `field:"optional" json:"managedCapacityConfiguration" yaml:"managedCapacityConfiguration"`
	// The maximum value that is allowed for the fleet's instance count for a location.
	//
	// When creating a new fleet, GameLift automatically sets this value to "1". Once the fleet is active, you can change this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/gamelift_fleet#max_size GameliftFleet#max_size}
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The minimum value allowed for the fleet's instance count for a location.
	//
	// When creating a new fleet, GameLift automatically sets this value to "0". After the fleet is active, you can change this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/gamelift_fleet#min_size GameliftFleet#min_size}
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
}

