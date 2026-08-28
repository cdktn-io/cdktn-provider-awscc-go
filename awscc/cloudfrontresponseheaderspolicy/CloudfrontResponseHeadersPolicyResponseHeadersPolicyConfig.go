// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontresponseheaderspolicy


type CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfig struct {
	// A name to identify the response headers policy.
	//
	// The name must be unique for response headers policies in this AWS-account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_response_headers_policy#name CloudfrontResponseHeadersPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A comment to describe the response headers policy.  The comment cannot be longer than 128 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_response_headers_policy#comment CloudfrontResponseHeadersPolicy#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// A configuration for a set of HTTP response headers that are used for cross-origin resource sharing (CORS).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_response_headers_policy#cors_config CloudfrontResponseHeadersPolicy#cors_config}
	CorsConfig *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfig `field:"optional" json:"corsConfig" yaml:"corsConfig"`
	// A configuration for a set of custom HTTP response headers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_response_headers_policy#custom_headers_config CloudfrontResponseHeadersPolicy#custom_headers_config}
	CustomHeadersConfig *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCustomHeadersConfig `field:"optional" json:"customHeadersConfig" yaml:"customHeadersConfig"`
	// A configuration for a set of HTTP headers to remove from the HTTP response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_response_headers_policy#remove_headers_config CloudfrontResponseHeadersPolicy#remove_headers_config}
	RemoveHeadersConfig *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigRemoveHeadersConfig `field:"optional" json:"removeHeadersConfig" yaml:"removeHeadersConfig"`
	// A configuration for a set of security-related HTTP response headers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_response_headers_policy#security_headers_config CloudfrontResponseHeadersPolicy#security_headers_config}
	SecurityHeadersConfig *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigSecurityHeadersConfig `field:"optional" json:"securityHeadersConfig" yaml:"securityHeadersConfig"`
	// A configuration for enabling the ``Server-Timing`` header in HTTP responses sent from CloudFront.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_response_headers_policy#server_timing_headers_config CloudfrontResponseHeadersPolicy#server_timing_headers_config}
	ServerTimingHeadersConfig *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigServerTimingHeadersConfig `field:"optional" json:"serverTimingHeadersConfig" yaml:"serverTimingHeadersConfig"`
}

