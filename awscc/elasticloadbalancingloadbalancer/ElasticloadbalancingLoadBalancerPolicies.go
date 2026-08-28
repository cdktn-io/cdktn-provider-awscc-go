// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingloadbalancer


type ElasticloadbalancingLoadBalancerPolicies struct {
	// The policy attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancing_load_balancer#attributes ElasticloadbalancingLoadBalancer#attributes}
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The instance ports for the policy. Required only for some policy types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancing_load_balancer#instance_ports ElasticloadbalancingLoadBalancer#instance_ports}
	InstancePorts *[]*string `field:"optional" json:"instancePorts" yaml:"instancePorts"`
	// The load balancer ports for the policy. Required only for some policy types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancing_load_balancer#load_balancer_ports ElasticloadbalancingLoadBalancer#load_balancer_ports}
	LoadBalancerPorts *[]*string `field:"optional" json:"loadBalancerPorts" yaml:"loadBalancerPorts"`
	// The name of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancing_load_balancer#policy_name ElasticloadbalancingLoadBalancer#policy_name}
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// The name of the policy type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancing_load_balancer#policy_type ElasticloadbalancingLoadBalancer#policy_type}
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
}

