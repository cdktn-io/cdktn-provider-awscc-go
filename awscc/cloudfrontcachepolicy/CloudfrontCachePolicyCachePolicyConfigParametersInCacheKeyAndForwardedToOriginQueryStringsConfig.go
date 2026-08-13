// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontcachepolicy


type CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginQueryStringsConfig struct {
	// Determines whether any URL query strings in viewer requests are included in the cache key and in requests that CloudFront sends to the origin.
	//
	// Valid values are:
	//   +  ``none`` ? No query strings in viewer requests are included in the cache key or in requests that CloudFront sends to the origin. Even when this field is set to ``none``, any query strings that are listed in an ``OriginRequestPolicy``*are* included in origin requests.
	//   +  ``whitelist`` ? Only the query strings in viewer requests that are listed in the ``QueryStringNames`` type are included in the cache key and in requests that CloudFront sends to the origin.
	//   +  ``allExcept`` ? All query strings in viewer requests are included in the cache key and in requests that CloudFront sends to the origin, *except* those that are listed in the ``QueryStringNames`` type, which are not included.
	//   +  ``all`` ? All query strings in viewer requests are included in the cache key and in requests that CloudFront sends to the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudfront_cache_policy#query_string_behavior CloudfrontCachePolicy#query_string_behavior}
	QueryStringBehavior *string `field:"required" json:"queryStringBehavior" yaml:"queryStringBehavior"`
	// Contains a list of query string names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudfront_cache_policy#query_strings CloudfrontCachePolicy#query_strings}
	QueryStrings *[]*string `field:"optional" json:"queryStrings" yaml:"queryStrings"`
}

