// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bitransformer


type B2BiTransformerEdiType struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/b2bi_transformer#x12_details B2BiTransformer#x12_details}.
	X12Details *B2BiTransformerEdiTypeX12Details `field:"optional" json:"x12Details" yaml:"x12Details"`
}

