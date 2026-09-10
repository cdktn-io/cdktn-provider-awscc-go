// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftfleet


type GameliftFleetPlayerGatewayConfiguration struct {
	// The IP protocol supported by the game server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/gamelift_fleet#game_server_ip_protocol_supported GameliftFleet#game_server_ip_protocol_supported}
	GameServerIpProtocolSupported *string `field:"optional" json:"gameServerIpProtocolSupported" yaml:"gameServerIpProtocolSupported"`
}

