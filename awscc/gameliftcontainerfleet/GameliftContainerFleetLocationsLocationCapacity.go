// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainerfleet


type GameliftContainerFleetLocationsLocationCapacity struct {
	// Defaults to MinSize if not defined.
	//
	// The number of EC2 instances you want to maintain in the specified fleet location. This value must fall between the minimum and maximum size limits. If any auto-scaling policy is defined for the container fleet, the desired instance will only be applied once during fleet creation and will be ignored in updates to avoid conflicts with auto-scaling. During updates with any auto-scaling policy defined, if current desired instance is lower than the new MinSize, it will be increased to the new MinSize; if current desired instance is larger than the new MaxSize, it will be decreased to the new MaxSize.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_fleet#desired_ec2_instances GameliftContainerFleet#desired_ec2_instances}
	DesiredEc2Instances *float64 `field:"optional" json:"desiredEc2Instances" yaml:"desiredEc2Instances"`
	// Configuration options for Amazon GameLift Servers-managed capacity behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_fleet#managed_capacity_configuration GameliftContainerFleet#managed_capacity_configuration}
	ManagedCapacityConfiguration *GameliftContainerFleetLocationsLocationCapacityManagedCapacityConfiguration `field:"optional" json:"managedCapacityConfiguration" yaml:"managedCapacityConfiguration"`
	// The maximum value that is allowed for the fleet's instance count for a location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_fleet#max_size GameliftContainerFleet#max_size}
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The minimum value allowed for the fleet's instance count for a location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_fleet#min_size GameliftContainerFleet#min_size}
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
}

