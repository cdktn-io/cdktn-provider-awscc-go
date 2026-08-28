// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingloadbalancer


type ElasticloadbalancingLoadBalancerHealthCheck struct {
	// The number of consecutive health checks successes required before moving the instance to the `Healthy` state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancing_load_balancer#healthy_threshold ElasticloadbalancingLoadBalancer#healthy_threshold}
	HealthyThreshold *string `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// The approximate interval, in seconds, between health checks of an individual instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancing_load_balancer#interval ElasticloadbalancingLoadBalancer#interval}
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
	// The instance being checked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancing_load_balancer#target ElasticloadbalancingLoadBalancer#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// The amount of time, in seconds, during which no response means a failed health check.
	//
	// This value must be less than the `Interval` value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancing_load_balancer#timeout ElasticloadbalancingLoadBalancer#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
	// The number of consecutive health check failures required before moving the instance to the `Unhealthy` state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancing_load_balancer#unhealthy_threshold ElasticloadbalancingLoadBalancer#unhealthy_threshold}
	UnhealthyThreshold *string `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

