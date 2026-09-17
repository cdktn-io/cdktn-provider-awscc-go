// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package omicsvariantstore


type OmicsVariantStoreSseConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/omics_variant_store#key_arn OmicsVariantStore#key_arn}.
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/omics_variant_store#type OmicsVariantStore#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

