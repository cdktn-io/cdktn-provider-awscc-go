// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncresolver


type AppsyncResolverCachingConfig struct {
	// The caching keys for a resolver that has caching activated.
	//
	// Valid values are entries from the ``$context.arguments``, ``$context.source``, and ``$context.identity`` maps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_resolver#caching_keys AppsyncResolver#caching_keys}
	CachingKeys *[]*string `field:"optional" json:"cachingKeys" yaml:"cachingKeys"`
	// The TTL in seconds for a resolver that has caching activated.  Valid values are 1?3,600 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_resolver#ttl AppsyncResolver#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

