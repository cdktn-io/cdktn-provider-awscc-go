// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesreceiptfilter


type SesReceiptFilterFilterIpFilter struct {
	// A single IP address or a range of IP addresses to block or allow, specified in CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ses_receipt_filter#cidr SesReceiptFilter#cidr}
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// Indicates whether to block or allow incoming mail from the specified IP addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ses_receipt_filter#policy SesReceiptFilter#policy}
	Policy *string `field:"required" json:"policy" yaml:"policy"`
}

