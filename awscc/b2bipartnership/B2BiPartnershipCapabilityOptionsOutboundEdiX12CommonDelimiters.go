// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bipartnership


type B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonDelimiters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/b2bi_partnership#component_separator B2BiPartnership#component_separator}.
	ComponentSeparator *string `field:"optional" json:"componentSeparator" yaml:"componentSeparator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/b2bi_partnership#data_element_separator B2BiPartnership#data_element_separator}.
	DataElementSeparator *string `field:"optional" json:"dataElementSeparator" yaml:"dataElementSeparator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/b2bi_partnership#segment_terminator B2BiPartnership#segment_terminator}.
	SegmentTerminator *string `field:"optional" json:"segmentTerminator" yaml:"segmentTerminator"`
}

