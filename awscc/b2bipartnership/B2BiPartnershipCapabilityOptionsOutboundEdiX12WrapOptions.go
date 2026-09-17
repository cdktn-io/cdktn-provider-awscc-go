// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bipartnership


type B2BiPartnershipCapabilityOptionsOutboundEdiX12WrapOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/b2bi_partnership#line_length B2BiPartnership#line_length}.
	LineLength *float64 `field:"optional" json:"lineLength" yaml:"lineLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/b2bi_partnership#line_terminator B2BiPartnership#line_terminator}.
	LineTerminator *string `field:"optional" json:"lineTerminator" yaml:"lineTerminator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/b2bi_partnership#wrap_by B2BiPartnership#wrap_by}.
	WrapBy *string `field:"optional" json:"wrapBy" yaml:"wrapBy"`
}

