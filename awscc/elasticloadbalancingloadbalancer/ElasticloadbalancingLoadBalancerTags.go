// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingloadbalancer


type ElasticloadbalancingLoadBalancerTags struct {
	// The key name of the tag.
	//
	// You can specify a value that's 1 to 128 Unicode characters in length and can't be prefixed with `aws:`. You can use any of the following characters: the set of Unicode letters, digits, whitespace, `_`, `.`, `/`, `=`, `+`, and `-`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticloadbalancing_load_balancer#key ElasticloadbalancingLoadBalancer#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag. You can specify a value that's 1 to 256 characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticloadbalancing_load_balancer#value ElasticloadbalancingLoadBalancer#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

