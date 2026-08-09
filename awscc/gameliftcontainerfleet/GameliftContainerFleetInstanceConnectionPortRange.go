// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainerfleet


type GameliftContainerFleetInstanceConnectionPortRange struct {
	// A starting value for a range of allowed port numbers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_fleet#from_port GameliftContainerFleet#from_port}
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// An ending value for a range of allowed port numbers.
	//
	// Port numbers are end-inclusive. This value must be higher than FromPort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_fleet#to_port GameliftContainerFleet#to_port}
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

