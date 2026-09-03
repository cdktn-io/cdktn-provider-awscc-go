// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bicapability


type B2BiCapabilityConfigurationEdiType struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/b2bi_capability#x12_details B2BiCapability#x12_details}.
	X12Details *B2BiCapabilityConfigurationEdiTypeX12Details `field:"optional" json:"x12Details" yaml:"x12Details"`
}

