// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagersitetositevpnattachment


type NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChange struct {
	// The rule number in the policy document that applies to this change.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_site_to_site_vpn_attachment#attachment_policy_rule_number NetworkmanagerSiteToSiteVpnAttachment#attachment_policy_rule_number}
	AttachmentPolicyRuleNumber *float64 `field:"optional" json:"attachmentPolicyRuleNumber" yaml:"attachmentPolicyRuleNumber"`
	// The name of the network function group to change.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_site_to_site_vpn_attachment#network_function_group_name NetworkmanagerSiteToSiteVpnAttachment#network_function_group_name}
	NetworkFunctionGroupName *string `field:"optional" json:"networkFunctionGroupName" yaml:"networkFunctionGroupName"`
	// The key-value tags that changed for the network function group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_site_to_site_vpn_attachment#tags NetworkmanagerSiteToSiteVpnAttachment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

