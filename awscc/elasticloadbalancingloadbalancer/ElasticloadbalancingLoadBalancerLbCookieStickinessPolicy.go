// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingloadbalancer


type ElasticloadbalancingLoadBalancerLbCookieStickinessPolicy struct {
	// The time period, in seconds, after which the cookie should be considered stale.
	//
	// If this parameter is not specified, the stickiness session lasts for the duration of the browser session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticloadbalancing_load_balancer#cookie_expiration_period ElasticloadbalancingLoadBalancer#cookie_expiration_period}
	CookieExpirationPeriod *string `field:"optional" json:"cookieExpirationPeriod" yaml:"cookieExpirationPeriod"`
	// The name of the policy. This name must be unique within the set of policies for this load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticloadbalancing_load_balancer#policy_name ElasticloadbalancingLoadBalancer#policy_name}
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

