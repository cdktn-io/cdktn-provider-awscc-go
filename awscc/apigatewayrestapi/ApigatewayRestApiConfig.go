// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayrestapi

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApigatewayRestApiConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#api_key_source_type ApigatewayRestApi#api_key_source_type}.
	ApiKeySourceType *string `field:"optional" json:"apiKeySourceType" yaml:"apiKeySourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#binary_media_types ApigatewayRestApi#binary_media_types}.
	BinaryMediaTypes *[]*string `field:"optional" json:"binaryMediaTypes" yaml:"binaryMediaTypes"`
	// An OpenAPI specification that defines a set of RESTful APIs in JSON format.
	//
	// For YAML templates, you can also provide the specification in YAML format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#body ApigatewayRestApi#body}
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The Amazon Simple Storage Service (Amazon S3) location that points to an OpenAPI file, which defines a set of RESTful APIs in JSON or YAML format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#body_s3_location ApigatewayRestApi#body_s3_location}
	BodyS3Location *ApigatewayRestApiBodyS3Location `field:"optional" json:"bodyS3Location" yaml:"bodyS3Location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#clone_from ApigatewayRestApi#clone_from}.
	CloneFrom *string `field:"optional" json:"cloneFrom" yaml:"cloneFrom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#description ApigatewayRestApi#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#disable_execute_api_endpoint ApigatewayRestApi#disable_execute_api_endpoint}.
	DisableExecuteApiEndpoint interface{} `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#endpoint_access_mode ApigatewayRestApi#endpoint_access_mode}.
	EndpointAccessMode *string `field:"optional" json:"endpointAccessMode" yaml:"endpointAccessMode"`
	// A list of the endpoint types and IP address types of the API.
	//
	// Use this property when creating an API. When importing an existing API, specify the endpoint configuration types using the ``Parameters`` property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#endpoint_configuration ApigatewayRestApi#endpoint_configuration}
	EndpointConfiguration *ApigatewayRestApiEndpointConfiguration `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#fail_on_warnings ApigatewayRestApi#fail_on_warnings}.
	FailOnWarnings interface{} `field:"optional" json:"failOnWarnings" yaml:"failOnWarnings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#minimum_compression_size ApigatewayRestApi#minimum_compression_size}.
	MinimumCompressionSize *float64 `field:"optional" json:"minimumCompressionSize" yaml:"minimumCompressionSize"`
	// This property applies only when you use OpenAPI to define your REST API.
	//
	// The ``Mode`` determines how API Gateway handles resource updates.
	//  Valid values are ``overwrite`` or ``merge``.
	//  For ``overwrite``, the new API definition replaces the existing one. The existing API identifier remains unchanged.
	//   For ``merge``, the new API definition is merged with the existing API.
	//  If you don't specify this property, a default value is chosen. For REST APIs created before March 29, 2021, the default is ``overwrite``. For REST APIs created after March 29, 2021, the new API definition takes precedence, but any container types such as endpoint configurations and binary media types are merged with the existing API.
	//  Use the default mode to define top-level ``RestApi`` properties in addition to using OpenAPI. Generally, it's preferred to use API Gateway's OpenAPI extensions to model these properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#mode ApigatewayRestApi#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The name of the RestApi.
	//
	// A name is required if the REST API is not based on an OpenAPI specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#name ApigatewayRestApi#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#parameters ApigatewayRestApi#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
	// A policy document that contains the permissions for the ``RestApi`` resource.
	//
	// To set the ARN for the policy, use the ``!Join`` intrinsic function with ``""`` as delimiter and values of ``"execute-api:/"`` and ``"*"``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#policy ApigatewayRestApi#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#security_policy ApigatewayRestApi#security_policy}.
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#tags ApigatewayRestApi#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#version ApigatewayRestApi#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

