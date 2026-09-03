// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagervpcattachment


type NetworkmanagerVpcAttachmentOptions struct {
	// Indicates whether to enable ApplianceModeSupport Support for Vpc Attachment. Valid Values: true | false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/networkmanager_vpc_attachment#appliance_mode_support NetworkmanagerVpcAttachment#appliance_mode_support}
	ApplianceModeSupport interface{} `field:"optional" json:"applianceModeSupport" yaml:"applianceModeSupport"`
	// Indicates whether to enable private DNS Support for Vpc Attachment. Valid Values: true | false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/networkmanager_vpc_attachment#dns_support NetworkmanagerVpcAttachment#dns_support}
	DnsSupport interface{} `field:"optional" json:"dnsSupport" yaml:"dnsSupport"`
	// Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/networkmanager_vpc_attachment#ipv_6_support NetworkmanagerVpcAttachment#ipv_6_support}
	Ipv6Support interface{} `field:"optional" json:"ipv6Support" yaml:"ipv6Support"`
	// Indicates whether to enable Security Group Referencing Support for Vpc Attachment. Valid Values: true | false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/networkmanager_vpc_attachment#security_group_referencing_support NetworkmanagerVpcAttachment#security_group_referencing_support}
	SecurityGroupReferencingSupport interface{} `field:"optional" json:"securityGroupReferencingSupport" yaml:"securityGroupReferencingSupport"`
}

