// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2loadbalancer


type Elasticloadbalancingv2LoadBalancerMinimumLoadBalancerCapacity struct {
	// The number of capacity units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_load_balancer#capacity_units Elasticloadbalancingv2LoadBalancer#capacity_units}
	CapacityUnits *float64 `field:"optional" json:"capacityUnits" yaml:"capacityUnits"`
}

