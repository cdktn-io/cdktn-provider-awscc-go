// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesreceiptfilter


type SesReceiptFilterFilter struct {
	// A structure that provides the IP addresses to block or allow, and whether to block or allow incoming mail from them.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_receipt_filter#ip_filter SesReceiptFilter#ip_filter}
	IpFilter *SesReceiptFilterFilterIpFilter `field:"required" json:"ipFilter" yaml:"ipFilter"`
	// The name of the IP address filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_receipt_filter#name SesReceiptFilter#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

