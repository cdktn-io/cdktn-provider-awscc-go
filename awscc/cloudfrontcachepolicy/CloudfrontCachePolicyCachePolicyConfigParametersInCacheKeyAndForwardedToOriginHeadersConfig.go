// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontcachepolicy


type CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfig struct {
	// Determines whether any HTTP headers are included in the cache key and in requests that CloudFront sends to the origin.
	//
	// Valid values are:
	//   +  ``none`` ? No HTTP headers are included in the cache key or in requests that CloudFront sends to the origin. Even when this field is set to ``none``, any headers that are listed in an ``OriginRequestPolicy``*are* included in origin requests.
	//   +  ``whitelist`` ? Only the HTTP headers that are listed in the ``Headers`` type are included in the cache key and in requests that CloudFront sends to the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudfront_cache_policy#header_behavior CloudfrontCachePolicy#header_behavior}
	HeaderBehavior *string `field:"required" json:"headerBehavior" yaml:"headerBehavior"`
	// Contains a list of HTTP header names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudfront_cache_policy#headers CloudfrontCachePolicy#headers}
	Headers *[]*string `field:"optional" json:"headers" yaml:"headers"`
}

