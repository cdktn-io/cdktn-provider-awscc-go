// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagersitetositevpnattachment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkmanagerSiteToSiteVpnAttachmentConfig struct {
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
	// The ID of a core network where you're creating a site-to-site VPN attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_site_to_site_vpn_attachment#core_network_id NetworkmanagerSiteToSiteVpnAttachment#core_network_id}
	CoreNetworkId *string `field:"required" json:"coreNetworkId" yaml:"coreNetworkId"`
	// The ARN of the site-to-site VPN attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_site_to_site_vpn_attachment#vpn_connection_arn NetworkmanagerSiteToSiteVpnAttachment#vpn_connection_arn}
	VpnConnectionArn *string `field:"required" json:"vpnConnectionArn" yaml:"vpnConnectionArn"`
	// The complete set of NAT gateway rules on the attachment.
	//
	// Updating this property replaces the entire rule set, and rules left out of the new value are removed. To remove all rules, specify an empty list or remove this property. NAT gateway support can't be enabled without rules, so it is then disabled unless NatGatewaySupport is specified. Only the core network owner can configure NAT gateway rules. The rules aren't returned when the attachment is read, so drift detection doesn't check them. Not available in AWS GovCloud (US).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_site_to_site_vpn_attachment#nat_gateway_rules NetworkmanagerSiteToSiteVpnAttachment#nat_gateway_rules}
	NatGatewayRules interface{} `field:"optional" json:"natGatewayRules" yaml:"natGatewayRules"`
	// Whether NAT gateway support is enabled on the attachment.
	//
	// Required when NatGatewayRules lists rules. Set to DISABLE with NatGatewayRules to keep the rules on the attachment without translating traffic, and set it to ENABLE to activate them again. An attachment that is pending acceptance applies this setting once it is accepted. Only the core network owner can configure or read NAT gateway settings. Not available in AWS GovCloud (US).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_site_to_site_vpn_attachment#nat_gateway_support NetworkmanagerSiteToSiteVpnAttachment#nat_gateway_support}
	NatGatewaySupport *string `field:"optional" json:"natGatewaySupport" yaml:"natGatewaySupport"`
	// Whether system routes are installed for the NAT CIDRs of the attachment's NAT gateway rules.
	//
	// Configurable on BGP Site-to-Site VPN attachments. Always false and not configurable for static VPN attachments. If omitted, a new NAT gateway configuration uses the service default for the attachment and an existing configuration keeps its current value, so removing this property from a template doesn't reset the setting. Only the core network owner can configure this setting. Not available in AWS GovCloud (US).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_site_to_site_vpn_attachment#nat_system_route_support NetworkmanagerSiteToSiteVpnAttachment#nat_system_route_support}
	NatSystemRouteSupport interface{} `field:"optional" json:"natSystemRouteSupport" yaml:"natSystemRouteSupport"`
	// The name of the network function group attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_site_to_site_vpn_attachment#network_function_group_name NetworkmanagerSiteToSiteVpnAttachment#network_function_group_name}
	NetworkFunctionGroupName *string `field:"optional" json:"networkFunctionGroupName" yaml:"networkFunctionGroupName"`
	// Routing policy label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_site_to_site_vpn_attachment#routing_policy_label NetworkmanagerSiteToSiteVpnAttachment#routing_policy_label}
	RoutingPolicyLabel *string `field:"optional" json:"routingPolicyLabel" yaml:"routingPolicyLabel"`
	// Tags for the attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmanager_site_to_site_vpn_attachment#tags NetworkmanagerSiteToSiteVpnAttachment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

