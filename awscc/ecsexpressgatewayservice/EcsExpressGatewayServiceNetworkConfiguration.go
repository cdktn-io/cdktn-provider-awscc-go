// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsexpressgatewayservice


type EcsExpressGatewayServiceNetworkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_express_gateway_service#security_groups EcsExpressGatewayService#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_express_gateway_service#subnets EcsExpressGatewayService#subnets}.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

