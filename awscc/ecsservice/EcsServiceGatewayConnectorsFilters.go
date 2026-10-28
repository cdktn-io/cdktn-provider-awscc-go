// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceGatewayConnectorsFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_service#access_types EcsService#access_types}.
	AccessTypes *[]*string `field:"optional" json:"accessTypes" yaml:"accessTypes"`
}

