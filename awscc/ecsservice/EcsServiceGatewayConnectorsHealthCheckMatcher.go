// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceGatewayConnectorsHealthCheckMatcher struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_service#grpc_code EcsService#grpc_code}.
	GrpcCode *string `field:"optional" json:"grpcCode" yaml:"grpcCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_service#http_code EcsService#http_code}.
	HttpCode *string `field:"optional" json:"httpCode" yaml:"httpCode"`
}

