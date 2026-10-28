// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsgateway


type EcsGatewayConfigurationHostSharedAlbAccessLogsDestinationS3 struct {
	// The S3 bucket for access logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#bucket EcsGateway#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 prefix for access logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#prefix EcsGateway#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

