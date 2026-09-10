// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabriclink


type RtbfabricLinkLinkAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rtbfabric_link#customer_provided_id RtbfabricLink#customer_provided_id}.
	CustomerProvidedId *string `field:"optional" json:"customerProvidedId" yaml:"customerProvidedId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rtbfabric_link#responder_error_masking RtbfabricLink#responder_error_masking}.
	ResponderErrorMasking interface{} `field:"optional" json:"responderErrorMasking" yaml:"responderErrorMasking"`
}

