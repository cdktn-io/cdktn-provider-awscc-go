// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package interconnectconnection


type InterconnectConnectionAttachPoint struct {
	// The ARN of the resource to attach to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/interconnect_connection#arn InterconnectConnection#arn}
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The ID of the Direct Connect Gateway to attach to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/interconnect_connection#direct_connect_gateway InterconnectConnection#direct_connect_gateway}
	DirectConnectGateway *string `field:"optional" json:"directConnectGateway" yaml:"directConnectGateway"`
}

