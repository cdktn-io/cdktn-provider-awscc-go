// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2api


type Apigatewayv2ApiBodyS3Location struct {
	// The S3 bucket that contains the OpenAPI definition to import. Required if you specify a ``BodyS3Location`` for an API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_api#bucket Apigatewayv2Api#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The Etag of the S3 object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_api#etag Apigatewayv2Api#etag}
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// The key of the S3 object. Required if you specify a ``BodyS3Location`` for an API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_api#key Apigatewayv2Api#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The version of the S3 object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_api#version Apigatewayv2Api#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

