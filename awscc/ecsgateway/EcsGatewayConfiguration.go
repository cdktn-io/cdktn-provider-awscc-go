// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsgateway


type EcsGatewayConfiguration struct {
	// The host shared ALB configuration for the gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#host_shared_alb EcsGateway#host_shared_alb}
	HostSharedAlb *EcsGatewayConfigurationHostSharedAlb `field:"required" json:"hostSharedAlb" yaml:"hostSharedAlb"`
}

