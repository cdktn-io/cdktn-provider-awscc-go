// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsgateway


type EcsGatewayConfigurationHostSharedAlbAccessLogsDestination struct {
	// The S3 destination for gateway ALB access logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#s3 EcsGateway#s3}
	S3 *EcsGatewayConfigurationHostSharedAlbAccessLogsDestinationS3 `field:"optional" json:"s3" yaml:"s3"`
}

