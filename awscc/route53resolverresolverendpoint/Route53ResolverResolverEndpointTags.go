// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53resolverresolverendpoint


type Route53ResolverResolverEndpointTags struct {
	// The name for the tag.
	//
	// For example, if you want to associate Resolver resources with the account IDs of your customers for billing purposes, the value of Key might be account-id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/route53resolver_resolver_endpoint#key Route53ResolverResolverEndpoint#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// For example, if Key is account-id, then Value might be the ID of the customer account that you're creating the resource for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/route53resolver_resolver_endpoint#value Route53ResolverResolverEndpoint#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

