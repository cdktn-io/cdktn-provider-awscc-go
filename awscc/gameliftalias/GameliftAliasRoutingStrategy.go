// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftalias


type GameliftAliasRoutingStrategy struct {
	// Simple routing strategy. The alias resolves to one specific fleet. Use this type when routing to active fleets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/gamelift_alias#type GameliftAlias#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A unique identifier for a fleet that the alias points to.
	//
	// If you specify SIMPLE for the Type property, you must specify this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/gamelift_alias#fleet_id GameliftAlias#fleet_id}
	FleetId *string `field:"optional" json:"fleetId" yaml:"fleetId"`
	// The message text to be used with a terminal routing strategy.
	//
	// If you specify TERMINAL for the Type property, you must specify this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/gamelift_alias#message GameliftAlias#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
}

