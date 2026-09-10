// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftfleet


type GameliftFleetAnywhereConfiguration struct {
	// Cost of compute can be specified on Anywhere Fleets to prioritize placement across Queue destinations based on Cost.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/gamelift_fleet#cost GameliftFleet#cost}
	Cost *string `field:"optional" json:"cost" yaml:"cost"`
}

