// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontcachepolicy


type CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOrigin struct {
	// An object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the cache key and in requests that CloudFront sends to the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudfront_cache_policy#cookies_config CloudfrontCachePolicy#cookies_config}
	CookiesConfig *CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginCookiesConfig `field:"required" json:"cookiesConfig" yaml:"cookiesConfig"`
	// A flag that can affect whether the ``Accept-Encoding`` HTTP header is included in the cache key and included in requests that CloudFront sends to the origin.
	//
	// This field is related to the ``EnableAcceptEncodingBrotli`` field. If one or both of these fields is ``true``*and* the viewer request includes the ``Accept-Encoding`` header, then CloudFront does the following:
	//   +  Normalizes the value of the viewer's ``Accept-Encoding`` header
	//   +  Includes the normalized header in the cache key
	//   +  Includes the normalized header in the request to the origin, if a request is necessary
	//
	//  For more information, see [Compression support](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-policy-compressed-objects) in the *Amazon CloudFront Developer Guide*.
	//  If you set this value to ``true``, and this cache behavior also has an origin request policy attached, do not include the ``Accept-Encoding`` header in the origin request policy. CloudFront always includes the ``Accept-Encoding`` header in origin requests when the value of this field is ``true``, so including this header in an origin request policy has no effect.
	//  If both of these fields are ``false``, then CloudFront treats the ``Accept-Encoding`` header the same as any other HTTP header in the viewer request. By default, it's not included in the cache key and it's not included in origin requests. In this case, you can manually add ``Accept-Encoding`` to the headers whitelist like any other HTTP header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudfront_cache_policy#enable_accept_encoding_gzip CloudfrontCachePolicy#enable_accept_encoding_gzip}
	EnableAcceptEncodingGzip interface{} `field:"required" json:"enableAcceptEncodingGzip" yaml:"enableAcceptEncodingGzip"`
	// An object that determines whether any HTTP headers (and if so, which headers) are included in the cache key and in requests that CloudFront sends to the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudfront_cache_policy#headers_config CloudfrontCachePolicy#headers_config}
	HeadersConfig *CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfig `field:"required" json:"headersConfig" yaml:"headersConfig"`
	// An object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the cache key and in requests that CloudFront sends to the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudfront_cache_policy#query_strings_config CloudfrontCachePolicy#query_strings_config}
	QueryStringsConfig *CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginQueryStringsConfig `field:"required" json:"queryStringsConfig" yaml:"queryStringsConfig"`
	// A flag that can affect whether the ``Accept-Encoding`` HTTP header is included in the cache key and included in requests that CloudFront sends to the origin.
	//
	// This field is related to the ``EnableAcceptEncodingGzip`` field. If one or both of these fields is ``true``*and* the viewer request includes the ``Accept-Encoding`` header, then CloudFront does the following:
	//   +  Normalizes the value of the viewer's ``Accept-Encoding`` header
	//   +  Includes the normalized header in the cache key
	//   +  Includes the normalized header in the request to the origin, if a request is necessary
	//
	//  For more information, see [Compression support](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-policy-compressed-objects) in the *Amazon CloudFront Developer Guide*.
	//  If you set this value to ``true``, and this cache behavior also has an origin request policy attached, do not include the ``Accept-Encoding`` header in the origin request policy. CloudFront always includes the ``Accept-Encoding`` header in origin requests when the value of this field is ``true``, so including this header in an origin request policy has no effect.
	//  If both of these fields are ``false``, then CloudFront treats the ``Accept-Encoding`` header the same as any other HTTP header in the viewer request. By default, it's not included in the cache key and it's not included in origin requests. In this case, you can manually add ``Accept-Encoding`` to the headers whitelist like any other HTTP header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cloudfront_cache_policy#enable_accept_encoding_brotli CloudfrontCachePolicy#enable_accept_encoding_brotli}
	EnableAcceptEncodingBrotli interface{} `field:"optional" json:"enableAcceptEncodingBrotli" yaml:"enableAcceptEncodingBrotli"`
}

