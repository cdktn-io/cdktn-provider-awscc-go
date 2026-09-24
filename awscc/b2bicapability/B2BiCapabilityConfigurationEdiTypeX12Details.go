// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bicapability


type B2BiCapabilityConfigurationEdiTypeX12Details struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/b2bi_capability#transaction_set B2BiCapability#transaction_set}.
	TransactionSet *string `field:"optional" json:"transactionSet" yaml:"transactionSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/b2bi_capability#version B2BiCapability#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

