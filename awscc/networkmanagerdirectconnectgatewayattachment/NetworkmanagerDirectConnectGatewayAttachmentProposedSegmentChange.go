// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagerdirectconnectgatewayattachment


type NetworkmanagerDirectConnectGatewayAttachmentProposedSegmentChange struct {
	// The rule number in the policy document that applies to this change.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_direct_connect_gateway_attachment#attachment_policy_rule_number NetworkmanagerDirectConnectGatewayAttachment#attachment_policy_rule_number}
	AttachmentPolicyRuleNumber *float64 `field:"optional" json:"attachmentPolicyRuleNumber" yaml:"attachmentPolicyRuleNumber"`
	// The name of the segment to change.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_direct_connect_gateway_attachment#segment_name NetworkmanagerDirectConnectGatewayAttachment#segment_name}
	SegmentName *string `field:"optional" json:"segmentName" yaml:"segmentName"`
	// The key-value tags that changed for the segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkmanager_direct_connect_gateway_attachment#tags NetworkmanagerDirectConnectGatewayAttachment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

