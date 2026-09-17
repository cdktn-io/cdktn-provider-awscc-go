// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2loadbalancer

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Elasticloadbalancingv2LoadBalancerConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Indicates whether to enable stabilization when creating or updating an LCU reservation.
	//
	// This ensures that the final stack status reflects the status of the LCU reservation. The default is ``false``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_load_balancer#enable_capacity_reservation_provision_stabilize Elasticloadbalancingv2LoadBalancer#enable_capacity_reservation_provision_stabilize}
	EnableCapacityReservationProvisionStabilize interface{} `field:"optional" json:"enableCapacityReservationProvisionStabilize" yaml:"enableCapacityReservationProvisionStabilize"`
	// [Network Load Balancers with UDP listeners] Indicates whether to use an IPv6 prefix from each subnet for source NAT.
	//
	// The IP address type must be ``dualstack``. The default value is ``off``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_load_balancer#enable_prefix_for_ipv_6_source_nat Elasticloadbalancingv2LoadBalancer#enable_prefix_for_ipv_6_source_nat}
	EnablePrefixForIpv6SourceNat *string `field:"optional" json:"enablePrefixForIpv6SourceNat" yaml:"enablePrefixForIpv6SourceNat"`
	// Indicates whether to evaluate inbound security group rules for traffic sent to a Network Load Balancer through privatelink.
	//
	// The default is ``on``.
	//  You can't configure this property on a Network Load Balancer unless you associated a security group with the load balancer when you created it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_load_balancer#enforce_security_group_inbound_rules_on_private_link_traffic Elasticloadbalancingv2LoadBalancer#enforce_security_group_inbound_rules_on_private_link_traffic}
	EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic *string `field:"optional" json:"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic" yaml:"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic"`
	// The IP address type.
	//
	// Internal load balancers must use ``ipv4``.
	//  [Application Load Balancers] The possible values are ``ipv4`` (IPv4 addresses), ``dualstack`` (IPv4 and IPv6 addresses), and ``dualstack-without-public-ipv4`` (public IPv6 addresses and private IPv4 and IPv6 addresses).
	//  Application Load Balancer authentication supports IPv4 addresses only when connecting to an Identity Provider (IdP) or Amazon Cognito endpoint. Without a public IPv4 address the load balancer can't complete the authentication process, resulting in HTTP 500 errors.
	//  [Network Load Balancers and Gateway Load Balancers] The possible values are ``ipv4`` (IPv4 addresses) and ``dualstack`` (IPv4 and IPv6 addresses).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_load_balancer#ip_address_type Elasticloadbalancingv2LoadBalancer#ip_address_type}
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// The ID of the IPv4 IPAM pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_load_balancer#ipv_4_ipam_pool_id Elasticloadbalancingv2LoadBalancer#ipv_4_ipam_pool_id}
	Ipv4IpamPoolId *string `field:"optional" json:"ipv4IpamPoolId" yaml:"ipv4IpamPoolId"`
	// The load balancer attributes. Attributes that you do not modify retain their current values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_load_balancer#load_balancer_attributes Elasticloadbalancingv2LoadBalancer#load_balancer_attributes}
	LoadBalancerAttributes interface{} `field:"optional" json:"loadBalancerAttributes" yaml:"loadBalancerAttributes"`
	// The minimum capacity for a load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_load_balancer#minimum_load_balancer_capacity Elasticloadbalancingv2LoadBalancer#minimum_load_balancer_capacity}
	MinimumLoadBalancerCapacity *Elasticloadbalancingv2LoadBalancerMinimumLoadBalancerCapacity `field:"optional" json:"minimumLoadBalancerCapacity" yaml:"minimumLoadBalancerCapacity"`
	// The name of the load balancer.
	//
	// This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, must not begin or end with a hyphen, and must not begin with "internal-".
	//  If you don't specify a name, AWS CloudFormation generates a unique physical ID for the load balancer. If you specify a name, you cannot perform updates that require replacement of this resource, but you can perform other updates. To replace the resource, specify a new name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_load_balancer#name Elasticloadbalancingv2LoadBalancer#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The nodes of an Internet-facing load balancer have public IP addresses.
	//
	// The DNS name of an Internet-facing load balancer is publicly resolvable to the public IP addresses of the nodes. Therefore, Internet-facing load balancers can route requests from clients over the internet.
	//  The nodes of an internal load balancer have only private IP addresses. The DNS name of an internal load balancer is publicly resolvable to the private IP addresses of the nodes. Therefore, internal load balancers can route requests only from clients with access to the VPC for the load balancer.
	//  The default is an Internet-facing load balancer.
	//  You can't specify a scheme for a Gateway Load Balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_load_balancer#scheme Elasticloadbalancingv2LoadBalancer#scheme}
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// [Application Load Balancers and Network Load Balancers] The IDs of the security groups for the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_load_balancer#security_groups Elasticloadbalancingv2LoadBalancer#security_groups}
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The IDs of the subnets.
	//
	// You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both.
	//  [Application Load Balancers] You must specify subnets from at least two Availability Zones. You can't specify Elastic IP addresses for your subnets.
	//  [Application Load Balancers on Outposts] You must specify one Outpost subnet.
	//  [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.
	//  [Network Load Balancers] You can specify subnets from one or more Availability Zones. You can specify one Elastic IP address per subnet if you need static IP addresses for your internet-facing load balancer. For internal load balancers, you can specify one private IP address per subnet from the IPv4 range of the subnet. For internet-facing load balancer, you can specify one IPv6 address per subnet.
	//  [Gateway Load Balancers] You can specify subnets from one or more Availability Zones. You can't specify Elastic IP addresses for your subnets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_load_balancer#subnet_mappings Elasticloadbalancingv2LoadBalancer#subnet_mappings}
	SubnetMappings interface{} `field:"optional" json:"subnetMappings" yaml:"subnetMappings"`
	// The IDs of the subnets.
	//
	// You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both. To specify an Elastic IP address, specify subnet mappings instead of subnets.
	//  [Application Load Balancers] You must specify subnets from at least two Availability Zones.
	//  [Application Load Balancers on Outposts] You must specify one Outpost subnet.
	//  [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.
	//  [Network Load Balancers and Gateway Load Balancers] You can specify subnets from one or more Availability Zones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_load_balancer#subnets Elasticloadbalancingv2LoadBalancer#subnets}
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
	// The tags to assign to the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_load_balancer#tags Elasticloadbalancingv2LoadBalancer#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The type of load balancer. The default is ``application``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticloadbalancingv2_load_balancer#type Elasticloadbalancingv2LoadBalancer#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

