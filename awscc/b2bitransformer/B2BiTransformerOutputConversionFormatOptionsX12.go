// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bitransformer


type B2BiTransformerOutputConversionFormatOptionsX12 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/b2bi_transformer#transaction_set B2BiTransformer#transaction_set}.
	TransactionSet *string `field:"optional" json:"transactionSet" yaml:"transactionSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/b2bi_transformer#version B2BiTransformer#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

