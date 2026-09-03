// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingloadbalancer


type ElasticloadbalancingLoadBalancerAppCookieStickinessPolicy struct {
	// The name of the application cookie used for stickiness.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/elasticloadbalancing_load_balancer#cookie_name ElasticloadbalancingLoadBalancer#cookie_name}
	CookieName *string `field:"optional" json:"cookieName" yaml:"cookieName"`
	// The mnemonic name for the policy being created.
	//
	// The name must be unique within a set of policies for this load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/elasticloadbalancing_load_balancer#policy_name ElasticloadbalancingLoadBalancer#policy_name}
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

