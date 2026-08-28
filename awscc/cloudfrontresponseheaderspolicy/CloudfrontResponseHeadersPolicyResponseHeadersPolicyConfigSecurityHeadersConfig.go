// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontresponseheaderspolicy


type CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigSecurityHeadersConfig struct {
	// The policy directives and their values that CloudFront includes as values for the ``Content-Security-Policy`` HTTP response header.
	//
	// For more information about the ``Content-Security-Policy`` HTTP response header, see [Content-Security-Policy](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy) in the MDN Web Docs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_response_headers_policy#content_security_policy CloudfrontResponseHeadersPolicy#content_security_policy}
	ContentSecurityPolicy *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigSecurityHeadersConfigContentSecurityPolicy `field:"optional" json:"contentSecurityPolicy" yaml:"contentSecurityPolicy"`
	// Determines whether CloudFront includes the ``X-Content-Type-Options`` HTTP response header with its value set to ``nosniff``.
	//
	// For more information about the ``X-Content-Type-Options`` HTTP response header, see [X-Content-Type-Options](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options) in the MDN Web Docs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_response_headers_policy#content_type_options CloudfrontResponseHeadersPolicy#content_type_options}
	ContentTypeOptions *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigSecurityHeadersConfigContentTypeOptions `field:"optional" json:"contentTypeOptions" yaml:"contentTypeOptions"`
	// Determines whether CloudFront includes the ``X-Frame-Options`` HTTP response header and the header's value.
	//
	// For more information about the ``X-Frame-Options`` HTTP response header, see [X-Frame-Options](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options) in the MDN Web Docs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_response_headers_policy#frame_options CloudfrontResponseHeadersPolicy#frame_options}
	FrameOptions *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigSecurityHeadersConfigFrameOptions `field:"optional" json:"frameOptions" yaml:"frameOptions"`
	// Determines whether CloudFront includes the ``Referrer-Policy`` HTTP response header and the header's value.
	//
	// For more information about the ``Referrer-Policy`` HTTP response header, see [Referrer-Policy](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy) in the MDN Web Docs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_response_headers_policy#referrer_policy CloudfrontResponseHeadersPolicy#referrer_policy}
	ReferrerPolicy *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigSecurityHeadersConfigReferrerPolicy `field:"optional" json:"referrerPolicy" yaml:"referrerPolicy"`
	// Determines whether CloudFront includes the ``Strict-Transport-Security`` HTTP response header and the header's value.
	//
	// For more information about the ``Strict-Transport-Security`` HTTP response header, see [Security headers](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/understanding-response-headers-policies.html#understanding-response-headers-policies-security) in the *Amazon CloudFront Developer Guide* and [Strict-Transport-Security](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security) in the MDN Web Docs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_response_headers_policy#strict_transport_security CloudfrontResponseHeadersPolicy#strict_transport_security}
	StrictTransportSecurity *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigSecurityHeadersConfigStrictTransportSecurity `field:"optional" json:"strictTransportSecurity" yaml:"strictTransportSecurity"`
	// Determines whether CloudFront includes the ``X-XSS-Protection`` HTTP response header and the header's value.
	//
	// For more information about the ``X-XSS-Protection`` HTTP response header, see [X-XSS-Protection](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_response_headers_policy#xss_protection CloudfrontResponseHeadersPolicy#xss_protection}
	XssProtection *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigSecurityHeadersConfigXssProtection `field:"optional" json:"xssProtection" yaml:"xssProtection"`
}

