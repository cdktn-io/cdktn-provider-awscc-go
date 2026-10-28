// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkfirewallfirewall

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkfirewallFirewallConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/networkfirewall_firewall#firewall_name NetworkfirewallFirewall#firewall_name}.
	FirewallName *string `field:"required" json:"firewallName" yaml:"firewallName"`
	// A resource ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/networkfirewall_firewall#firewall_policy_arn NetworkfirewallFirewall#firewall_policy_arn}
	FirewallPolicyArn *string `field:"required" json:"firewallPolicyArn" yaml:"firewallPolicyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/networkfirewall_firewall#availability_zone_change_protection NetworkfirewallFirewall#availability_zone_change_protection}.
	AvailabilityZoneChangeProtection interface{} `field:"optional" json:"availabilityZoneChangeProtection" yaml:"availabilityZoneChangeProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/networkfirewall_firewall#availability_zone_mappings NetworkfirewallFirewall#availability_zone_mappings}.
	AvailabilityZoneMappings interface{} `field:"optional" json:"availabilityZoneMappings" yaml:"availabilityZoneMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/networkfirewall_firewall#delete_protection NetworkfirewallFirewall#delete_protection}.
	DeleteProtection interface{} `field:"optional" json:"deleteProtection" yaml:"deleteProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/networkfirewall_firewall#description NetworkfirewallFirewall#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The types of analysis to enable for the firewall. Can be TLS_SNI, HTTP_HOST, or both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/networkfirewall_firewall#enabled_analysis_types NetworkfirewallFirewall#enabled_analysis_types}
	EnabledAnalysisTypes *[]*string `field:"optional" json:"enabledAnalysisTypes" yaml:"enabledAnalysisTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/networkfirewall_firewall#firewall_policy_change_protection NetworkfirewallFirewall#firewall_policy_change_protection}.
	FirewallPolicyChangeProtection interface{} `field:"optional" json:"firewallPolicyChangeProtection" yaml:"firewallPolicyChangeProtection"`
	// The NAT gateways that the firewall proxies egress traffic through when NoSourcePreservation is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/networkfirewall_firewall#nat_gateway_mappings NetworkfirewallFirewall#nat_gateway_mappings}
	NatGatewayMappings interface{} `field:"optional" json:"natGatewayMappings" yaml:"natGatewayMappings"`
	// Indicates whether the firewall operates in proxy mode, where the source IP address of traffic is not preserved.
	//
	// When true, specify NatGatewayMappings and VpcEndpoint instead of the top-level VpcId and SubnetMappings. You can't change this setting after you create the firewall.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/networkfirewall_firewall#no_source_preservation NetworkfirewallFirewall#no_source_preservation}
	NoSourcePreservation interface{} `field:"optional" json:"noSourcePreservation" yaml:"noSourcePreservation"`
	// The proxy listener configuration for a proxy-mode (no source preservation) firewall.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/networkfirewall_firewall#proxy_settings NetworkfirewallFirewall#proxy_settings}
	ProxySettings *NetworkfirewallFirewallProxySettings `field:"optional" json:"proxySettings" yaml:"proxySettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/networkfirewall_firewall#subnet_change_protection NetworkfirewallFirewall#subnet_change_protection}.
	SubnetChangeProtection interface{} `field:"optional" json:"subnetChangeProtection" yaml:"subnetChangeProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/networkfirewall_firewall#subnet_mappings NetworkfirewallFirewall#subnet_mappings}.
	SubnetMappings interface{} `field:"optional" json:"subnetMappings" yaml:"subnetMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/networkfirewall_firewall#tags NetworkfirewallFirewall#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/networkfirewall_firewall#transit_gateway_id NetworkfirewallFirewall#transit_gateway_id}.
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The VPC and subnets for a proxy-mode firewall's endpoint.
	//
	// Provide this instead of the top-level VpcId and SubnetMappings when NoSourcePreservation is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/networkfirewall_firewall#vpc_endpoint NetworkfirewallFirewall#vpc_endpoint}
	VpcEndpoint *NetworkfirewallFirewallVpcEndpoint `field:"optional" json:"vpcEndpoint" yaml:"vpcEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/networkfirewall_firewall#vpc_id NetworkfirewallFirewall#vpc_id}.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

