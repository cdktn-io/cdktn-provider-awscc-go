// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bipartnership


type B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/b2bi_partnership#starting_functional_group_control_number B2BiPartnership#starting_functional_group_control_number}.
	StartingFunctionalGroupControlNumber *float64 `field:"optional" json:"startingFunctionalGroupControlNumber" yaml:"startingFunctionalGroupControlNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/b2bi_partnership#starting_interchange_control_number B2BiPartnership#starting_interchange_control_number}.
	StartingInterchangeControlNumber *float64 `field:"optional" json:"startingInterchangeControlNumber" yaml:"startingInterchangeControlNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/b2bi_partnership#starting_transaction_set_control_number B2BiPartnership#starting_transaction_set_control_number}.
	StartingTransactionSetControlNumber *float64 `field:"optional" json:"startingTransactionSetControlNumber" yaml:"startingTransactionSetControlNumber"`
}

