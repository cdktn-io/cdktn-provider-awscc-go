// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayrestapi


type ApigatewayRestApiBodyS3Location struct {
	// The name of the S3 bucket where the OpenAPI file is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apigateway_rest_api#bucket ApigatewayRestApi#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The Amazon S3 ETag (a file checksum) of the OpenAPI file.
	//
	// If you don't specify a value, API Gateway skips ETag validation of your OpenAPI file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apigateway_rest_api#e_tag ApigatewayRestApi#e_tag}
	ETag *string `field:"optional" json:"eTag" yaml:"eTag"`
	// The file name of the OpenAPI file (Amazon S3 object name).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apigateway_rest_api#key ApigatewayRestApi#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// For versioning-enabled buckets, a specific version of the OpenAPI file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apigateway_rest_api#version ApigatewayRestApi#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

