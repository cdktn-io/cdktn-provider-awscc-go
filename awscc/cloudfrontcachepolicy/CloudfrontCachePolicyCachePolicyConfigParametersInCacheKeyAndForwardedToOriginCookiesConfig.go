// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontcachepolicy


type CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginCookiesConfig struct {
	// Determines whether any cookies in viewer requests are included in the cache key and in requests that CloudFront sends to the origin.
	//
	// Valid values are:
	//   +  ``none`` ? No cookies in viewer requests are included in the cache key or in requests that CloudFront sends to the origin. Even when this field is set to ``none``, any cookies that are listed in an ``OriginRequestPolicy``*are* included in origin requests.
	//   +  ``whitelist`` ? Only the cookies in viewer requests that are listed in the ``CookieNames`` type are included in the cache key and in requests that CloudFront sends to the origin.
	//   +  ``allExcept`` ? All cookies in viewer requests are included in the cache key and in requests that CloudFront sends to the origin, *except* for those that are listed in the ``CookieNames`` type, which are not included.
	//   +  ``all`` ? All cookies in viewer requests are included in the cache key and in requests that CloudFront sends to the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_cache_policy#cookie_behavior CloudfrontCachePolicy#cookie_behavior}
	CookieBehavior *string `field:"required" json:"cookieBehavior" yaml:"cookieBehavior"`
	// Contains a list of cookie names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_cache_policy#cookies CloudfrontCachePolicy#cookies}
	Cookies *[]*string `field:"optional" json:"cookies" yaml:"cookies"`
}

