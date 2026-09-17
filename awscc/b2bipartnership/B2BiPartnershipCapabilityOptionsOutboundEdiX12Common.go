// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bipartnership


type B2BiPartnershipCapabilityOptionsOutboundEdiX12Common struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/b2bi_partnership#control_numbers B2BiPartnership#control_numbers}.
	ControlNumbers *B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonControlNumbers `field:"optional" json:"controlNumbers" yaml:"controlNumbers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/b2bi_partnership#delimiters B2BiPartnership#delimiters}.
	Delimiters *B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonDelimiters `field:"optional" json:"delimiters" yaml:"delimiters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/b2bi_partnership#functional_group_headers B2BiPartnership#functional_group_headers}.
	FunctionalGroupHeaders *B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonFunctionalGroupHeaders `field:"optional" json:"functionalGroupHeaders" yaml:"functionalGroupHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/b2bi_partnership#gs_05_time_format B2BiPartnership#gs_05_time_format}.
	Gs05TimeFormat *string `field:"optional" json:"gs05TimeFormat" yaml:"gs05TimeFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/b2bi_partnership#interchange_control_headers B2BiPartnership#interchange_control_headers}.
	InterchangeControlHeaders *B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonInterchangeControlHeaders `field:"optional" json:"interchangeControlHeaders" yaml:"interchangeControlHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/b2bi_partnership#validate_edi B2BiPartnership#validate_edi}.
	ValidateEdi interface{} `field:"optional" json:"validateEdi" yaml:"validateEdi"`
}

