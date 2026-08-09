// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagerconnectattachment


type NetworkmanagerConnectAttachmentOptions struct {
	// Tunnel protocol for connect attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/networkmanager_connect_attachment#protocol NetworkmanagerConnectAttachment#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

