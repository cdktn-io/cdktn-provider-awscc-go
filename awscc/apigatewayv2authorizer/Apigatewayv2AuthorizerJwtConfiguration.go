// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2authorizer


type Apigatewayv2AuthorizerJwtConfiguration struct {
	// A list of the intended recipients of the JWT.
	//
	// A valid JWT must provide an ``aud`` that matches at least one entry in this list. See [RFC 7519](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc7519#section-4.1.3). Required for the ``JWT`` authorizer type. Supported only for HTTP APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_authorizer#audience Apigatewayv2Authorizer#audience}
	Audience *[]*string `field:"optional" json:"audience" yaml:"audience"`
	// The base domain of the identity provider that issues JSON Web Tokens.
	//
	// For example, an Amazon Cognito user pool has the following format: ``https://cognito-idp.{region}.amazonaws.com/{userPoolId}``. Required for the ``JWT`` authorizer type. Supported only for HTTP APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_authorizer#issuer Apigatewayv2Authorizer#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
}

