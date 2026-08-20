// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2api

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Apigatewayv2ApiConfig struct {
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
	// An API key selection expression. Supported only for WebSocket APIs. See [API Key Selection Expressions](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api#api_key_selection_expression Apigatewayv2Api#api_key_selection_expression}
	ApiKeySelectionExpression *string `field:"optional" json:"apiKeySelectionExpression" yaml:"apiKeySelectionExpression"`
	// Specifies how to interpret the base path of the API during import.
	//
	// Valid values are ``ignore``, ``prepend``, and ``split``. The default value is ``ignore``. To learn more, see [Set the OpenAPI basePath Property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api-basePath.html). Supported only for HTTP APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api#base_path Apigatewayv2Api#base_path}
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// The OpenAPI definition.
	//
	// Supported only for HTTP APIs. To import an HTTP API, you must specify a ``Body`` or ``BodyS3Location``. If you specify a ``Body`` or ``BodyS3Location``, don't specify CloudFormation resources such as ``AWS::ApiGatewayV2::Authorizer`` or ``AWS::ApiGatewayV2::Route``. API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api#body Apigatewayv2Api#body}
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The S3 location of an OpenAPI definition.
	//
	// Supported only for HTTP APIs. To import an HTTP API, you must specify a ``Body`` or ``BodyS3Location``. If you specify a ``Body`` or ``BodyS3Location``, don't specify CloudFormation resources such as ``AWS::ApiGatewayV2::Authorizer`` or ``AWS::ApiGatewayV2::Route``. API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api#body_s3_location Apigatewayv2Api#body_s3_location}
	BodyS3Location *Apigatewayv2ApiBodyS3Location `field:"optional" json:"bodyS3Location" yaml:"bodyS3Location"`
	// A CORS configuration. Supported only for HTTP APIs. See [Configuring CORS](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html) for more information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api#cors_configuration Apigatewayv2Api#cors_configuration}
	CorsConfiguration *Apigatewayv2ApiCorsConfiguration `field:"optional" json:"corsConfiguration" yaml:"corsConfiguration"`
	// This property is part of quick create.
	//
	// It specifies the credentials required for the integration, if any. For a Lambda integration, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify ``arn:aws:iam::*:user/*``. To use resource-based permissions on supported AWS services, specify ``null``. Currently, this property is not used for HTTP integrations. Supported only for HTTP APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api#credentials_arn Apigatewayv2Api#credentials_arn}
	CredentialsArn *string `field:"optional" json:"credentialsArn" yaml:"credentialsArn"`
	// The description of the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api#description Apigatewayv2Api#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether clients can invoke your API by using the default ``execute-api`` endpoint.
	//
	// By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api#disable_execute_api_endpoint Apigatewayv2Api#disable_execute_api_endpoint}
	DisableExecuteApiEndpoint interface{} `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// Avoid validating models when creating a deployment. Supported only for WebSocket APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api#disable_schema_validation Apigatewayv2Api#disable_schema_validation}
	DisableSchemaValidation interface{} `field:"optional" json:"disableSchemaValidation" yaml:"disableSchemaValidation"`
	// Specifies whether to rollback the API creation when a warning is encountered.
	//
	// By default, API creation continues if a warning is encountered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api#fail_on_warnings Apigatewayv2Api#fail_on_warnings}
	FailOnWarnings interface{} `field:"optional" json:"failOnWarnings" yaml:"failOnWarnings"`
	// The IP address types that can invoke the API.
	//
	// Use ``ipv4`` to allow only IPv4 addresses to invoke your API, or use ``dualstack`` to allow both IPv4 and IPv6 addresses to invoke your API.
	//  Don?t use IP address type for an HTTP API based on an OpenAPI specification. Instead, specify the IP address type in the OpenAPI specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api#ip_address_type Apigatewayv2Api#ip_address_type}
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// The name of the API. Required unless you specify an OpenAPI definition for ``Body`` or ``S3BodyLocation``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api#name Apigatewayv2Api#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The API protocol.
	//
	// Valid values are ``WEBSOCKET`` or ``HTTP``. Required unless you specify an OpenAPI definition for ``Body`` or ``S3BodyLocation``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api#protocol_type Apigatewayv2Api#protocol_type}
	ProtocolType *string `field:"optional" json:"protocolType" yaml:"protocolType"`
	// This property is part of quick create.
	//
	// If you don't specify a ``routeKey``, a default route of ``$default`` is created. The ``$default`` route acts as a catch-all for any request made to your API, for a particular stage. The ``$default`` route key can't be modified. You can add routes after creating the API, and you can update the route keys of additional routes. Supported only for HTTP APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api#route_key Apigatewayv2Api#route_key}
	RouteKey *string `field:"optional" json:"routeKey" yaml:"routeKey"`
	// The route selection expression for the API.
	//
	// For HTTP APIs, the ``routeSelectionExpression`` must be ``${request.method} ${request.path}``. If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api#route_selection_expression Apigatewayv2Api#route_selection_expression}
	RouteSelectionExpression *string `field:"optional" json:"routeSelectionExpression" yaml:"routeSelectionExpression"`
	// The collection of tags. Each tag element is associated with a given resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api#tags Apigatewayv2Api#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// This property is part of quick create.
	//
	// Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes. For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN. The type of the integration will be HTTP_PROXY or AWS_PROXY, respectively. Supported only for HTTP APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api#target Apigatewayv2Api#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// A version identifier for the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api#version Apigatewayv2Api#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

