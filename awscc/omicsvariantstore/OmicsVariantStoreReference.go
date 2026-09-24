// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package omicsvariantstore


type OmicsVariantStoreReference struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/omics_variant_store#reference_arn OmicsVariantStore#reference_arn}.
	ReferenceArn *string `field:"required" json:"referenceArn" yaml:"referenceArn"`
}

