// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainerfleet


type GameliftContainerFleetInstanceInboundPermissions struct {
	// A starting value for a range of allowed port numbers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/gamelift_container_fleet#from_port GameliftContainerFleet#from_port}
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// A range of allowed IP addresses.
	//
	// This value must be expressed in CIDR notation. Example: "000.000.000.000/[subnet mask]" or optionally the shortened version "0.0.0.0/[subnet mask]".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/gamelift_container_fleet#ip_range GameliftContainerFleet#ip_range}
	IpRange *string `field:"optional" json:"ipRange" yaml:"ipRange"`
	// The network communication protocol used by the fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/gamelift_container_fleet#protocol GameliftContainerFleet#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// An ending value for a range of allowed port numbers.
	//
	// Port numbers are end-inclusive. This value must be higher than FromPort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/gamelift_container_fleet#to_port GameliftContainerFleet#to_port}
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

