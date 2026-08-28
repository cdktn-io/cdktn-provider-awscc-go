// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2api


type Apigatewayv2ApiCorsConfiguration struct {
	// Specifies whether credentials are included in the CORS request. Supported only for HTTP APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_api#allow_credentials Apigatewayv2Api#allow_credentials}
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// Represents a collection of allowed headers. Supported only for HTTP APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_api#allow_headers Apigatewayv2Api#allow_headers}
	AllowHeaders *[]*string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// Represents a collection of allowed HTTP methods. Supported only for HTTP APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_api#allow_methods Apigatewayv2Api#allow_methods}
	AllowMethods *[]*string `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// Represents a collection of allowed origins. Supported only for HTTP APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_api#allow_origins Apigatewayv2Api#allow_origins}
	AllowOrigins *[]*string `field:"optional" json:"allowOrigins" yaml:"allowOrigins"`
	// Represents a collection of exposed headers. Supported only for HTTP APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_api#expose_headers Apigatewayv2Api#expose_headers}
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
	// The number of seconds that the browser should cache preflight request results. Supported only for HTTP APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_api#max_age Apigatewayv2Api#max_age}
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
}

