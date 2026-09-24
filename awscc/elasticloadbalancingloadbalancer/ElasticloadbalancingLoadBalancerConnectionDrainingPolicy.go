// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingloadbalancer


type ElasticloadbalancingLoadBalancerConnectionDrainingPolicy struct {
	// Specifies whether connection draining is enabled for the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/elasticloadbalancing_load_balancer#enabled ElasticloadbalancingLoadBalancer#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The maximum time, in seconds, to keep the existing connections open before deregistering the instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/elasticloadbalancing_load_balancer#timeout ElasticloadbalancingLoadBalancer#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

