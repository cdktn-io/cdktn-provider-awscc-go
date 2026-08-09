// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontoriginrequestpolicy


type CloudfrontOriginRequestPolicyOriginRequestPolicyConfig struct {
	// The cookies from viewer requests to include in origin requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudfront_origin_request_policy#cookies_config CloudfrontOriginRequestPolicy#cookies_config}
	CookiesConfig *CloudfrontOriginRequestPolicyOriginRequestPolicyConfigCookiesConfig `field:"required" json:"cookiesConfig" yaml:"cookiesConfig"`
	// The HTTP headers to include in origin requests.
	//
	// These can include headers from viewer requests and additional headers added by CloudFront.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudfront_origin_request_policy#headers_config CloudfrontOriginRequestPolicy#headers_config}
	HeadersConfig *CloudfrontOriginRequestPolicyOriginRequestPolicyConfigHeadersConfig `field:"required" json:"headersConfig" yaml:"headersConfig"`
	// A unique name to identify the origin request policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudfront_origin_request_policy#name CloudfrontOriginRequestPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The URL query strings from viewer requests to include in origin requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudfront_origin_request_policy#query_strings_config CloudfrontOriginRequestPolicy#query_strings_config}
	QueryStringsConfig *CloudfrontOriginRequestPolicyOriginRequestPolicyConfigQueryStringsConfig `field:"required" json:"queryStringsConfig" yaml:"queryStringsConfig"`
	// A comment to describe the origin request policy. The comment cannot be longer than 128 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudfront_origin_request_policy#comment CloudfrontOriginRequestPolicy#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

