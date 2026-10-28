// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsgateway


type EcsGatewayConfigurationHostSharedAlbAccessLogs struct {
	// The destination for gateway ALB access logs. Exactly one destination type must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#destination EcsGateway#destination}
	Destination *EcsGatewayConfigurationHostSharedAlbAccessLogsDestination `field:"optional" json:"destination" yaml:"destination"`
}

