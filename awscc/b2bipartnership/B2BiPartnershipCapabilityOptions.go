// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bipartnership


type B2BiPartnershipCapabilityOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/b2bi_partnership#inbound_edi B2BiPartnership#inbound_edi}.
	InboundEdi *B2BiPartnershipCapabilityOptionsInboundEdi `field:"optional" json:"inboundEdi" yaml:"inboundEdi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/b2bi_partnership#outbound_edi B2BiPartnership#outbound_edi}.
	OutboundEdi *B2BiPartnershipCapabilityOptionsOutboundEdi `field:"optional" json:"outboundEdi" yaml:"outboundEdi"`
}

