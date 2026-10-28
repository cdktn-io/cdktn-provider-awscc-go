// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsgateway


type EcsGatewayConfigurationHostSharedAlbTestTrafficRulesHeader struct {
	// The name of the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#name EcsGateway#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The values for the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#values EcsGateway#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

