// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bipartnership


type B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonFunctionalGroupHeaders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/b2bi_partnership#application_receiver_code B2BiPartnership#application_receiver_code}.
	ApplicationReceiverCode *string `field:"optional" json:"applicationReceiverCode" yaml:"applicationReceiverCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/b2bi_partnership#application_sender_code B2BiPartnership#application_sender_code}.
	ApplicationSenderCode *string `field:"optional" json:"applicationSenderCode" yaml:"applicationSenderCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/b2bi_partnership#responsible_agency_code B2BiPartnership#responsible_agency_code}.
	ResponsibleAgencyCode *string `field:"optional" json:"responsibleAgencyCode" yaml:"responsibleAgencyCode"`
}

