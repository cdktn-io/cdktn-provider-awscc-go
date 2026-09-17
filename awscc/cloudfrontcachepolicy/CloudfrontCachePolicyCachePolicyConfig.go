// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontcachepolicy


type CloudfrontCachePolicyCachePolicyConfig struct {
	// The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	//
	// CloudFront uses this value as the object's time to live (TTL) only when the origin does *not* send ``Cache-Control`` or ``Expires`` headers with the object. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide*.
	//  The default value for this field is 86400 seconds (one day). If the value of ``MinTTL`` is more than 86400 seconds, then the default value for this field is the same as the value of ``MinTTL``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudfront_cache_policy#default_ttl CloudfrontCachePolicy#default_ttl}
	DefaultTtl *float64 `field:"required" json:"defaultTtl" yaml:"defaultTtl"`
	// The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	//
	// CloudFront uses this value only when the origin sends ``Cache-Control`` or ``Expires`` headers with the object. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide*.
	//  The default value for this field is 31536000 seconds (one year). If the value of ``MinTTL`` or ``DefaultTTL`` is more than 31536000 seconds, then the default value for this field is the same as the value of ``DefaultTTL``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudfront_cache_policy#max_ttl CloudfrontCachePolicy#max_ttl}
	MaxTtl *float64 `field:"required" json:"maxTtl" yaml:"maxTtl"`
	// The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	//
	// For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudfront_cache_policy#min_ttl CloudfrontCachePolicy#min_ttl}
	MinTtl *float64 `field:"required" json:"minTtl" yaml:"minTtl"`
	// A unique name to identify the cache policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudfront_cache_policy#name CloudfrontCachePolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The HTTP headers, cookies, and URL query strings to include in the cache key.
	//
	// The values included in the cache key are also included in requests that CloudFront sends to the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudfront_cache_policy#parameters_in_cache_key_and_forwarded_to_origin CloudfrontCachePolicy#parameters_in_cache_key_and_forwarded_to_origin}
	ParametersInCacheKeyAndForwardedToOrigin *CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOrigin `field:"required" json:"parametersInCacheKeyAndForwardedToOrigin" yaml:"parametersInCacheKeyAndForwardedToOrigin"`
	// A comment to describe the cache policy. The comment cannot be longer than 128 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudfront_cache_policy#comment CloudfrontCachePolicy#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

