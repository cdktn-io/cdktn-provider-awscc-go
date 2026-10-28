// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceGatewayConnectors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_service#filters EcsService#filters}.
	Filters *EcsServiceGatewayConnectorsFilters `field:"optional" json:"filters" yaml:"filters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_service#health_check EcsService#health_check}.
	HealthCheck *EcsServiceGatewayConnectorsHealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_service#port_name EcsService#port_name}.
	PortName *string `field:"optional" json:"portName" yaml:"portName"`
}

