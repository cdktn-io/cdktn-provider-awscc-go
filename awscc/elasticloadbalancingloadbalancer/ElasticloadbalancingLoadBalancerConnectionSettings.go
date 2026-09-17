// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingloadbalancer


type ElasticloadbalancingLoadBalancerConnectionSettings struct {
	// The time, in seconds, that the connection is allowed to be idle (no data has been sent over the connection) before it is closed by the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancing_load_balancer#idle_timeout ElasticloadbalancingLoadBalancer#idle_timeout}
	IdleTimeout *float64 `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
}

