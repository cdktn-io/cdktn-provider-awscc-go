// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorfunction


type MediatailorFunctionHttpRequestConfigurationCache struct {
	// A JSONata expression that MediaTailor evaluates to a custom cache key.
	//
	// By default, the cache key is a hash of the HTTP URL, request body, and HTTP method; request headers are not included. The evaluated key must be smaller than 1 KB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediatailor_function#key MediatailorFunction#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The upper bound, in seconds, on how long MediaTailor caches a response.
	//
	// This value must be greater than or equal to TtlMinimumSeconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediatailor_function#ttl_maximum_seconds MediatailorFunction#ttl_maximum_seconds}
	TtlMaximumSeconds *float64 `field:"optional" json:"ttlMaximumSeconds" yaml:"ttlMaximumSeconds"`
	// The lower bound, in seconds, on how long MediaTailor caches a response.
	//
	// MediaTailor also uses this value as the cache duration when a response has no Cache-Control max-age directive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediatailor_function#ttl_minimum_seconds MediatailorFunction#ttl_minimum_seconds}
	TtlMinimumSeconds *float64 `field:"optional" json:"ttlMinimumSeconds" yaml:"ttlMinimumSeconds"`
}

