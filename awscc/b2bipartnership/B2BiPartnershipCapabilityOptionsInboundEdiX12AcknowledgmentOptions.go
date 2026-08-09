// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bipartnership


type B2BiPartnershipCapabilityOptionsInboundEdiX12AcknowledgmentOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/b2bi_partnership#functional_acknowledgment B2BiPartnership#functional_acknowledgment}.
	FunctionalAcknowledgment *string `field:"optional" json:"functionalAcknowledgment" yaml:"functionalAcknowledgment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/b2bi_partnership#technical_acknowledgment B2BiPartnership#technical_acknowledgment}.
	TechnicalAcknowledgment *string `field:"optional" json:"technicalAcknowledgment" yaml:"technicalAcknowledgment"`
}

