// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsgateway


type EcsGatewayConfigurationHostSharedAlbTestTrafficRules struct {
	// The header configuration for test traffic routing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#header EcsGateway#header}
	Header *EcsGatewayConfigurationHostSharedAlbTestTrafficRulesHeader `field:"optional" json:"header" yaml:"header"`
}

