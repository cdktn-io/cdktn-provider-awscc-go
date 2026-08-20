// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bipartnership


type B2BiPartnershipCapabilityOptionsOutboundEdiX12 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/b2bi_partnership#common B2BiPartnership#common}.
	Common *B2BiPartnershipCapabilityOptionsOutboundEdiX12Common `field:"optional" json:"common" yaml:"common"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/b2bi_partnership#wrap_options B2BiPartnership#wrap_options}.
	WrapOptions *B2BiPartnershipCapabilityOptionsOutboundEdiX12WrapOptions `field:"optional" json:"wrapOptions" yaml:"wrapOptions"`
}

