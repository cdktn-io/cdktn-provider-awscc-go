// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabricoutboundexternallink


type RtbfabricOutboundExternalLinkLinkAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_outbound_external_link#customer_provided_id RtbfabricOutboundExternalLink#customer_provided_id}.
	CustomerProvidedId *string `field:"optional" json:"customerProvidedId" yaml:"customerProvidedId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_outbound_external_link#responder_error_masking RtbfabricOutboundExternalLink#responder_error_masking}.
	ResponderErrorMasking interface{} `field:"optional" json:"responderErrorMasking" yaml:"responderErrorMasking"`
}

