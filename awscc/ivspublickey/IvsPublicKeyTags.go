// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ivspublickey


type IvsPublicKeyTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ivs_public_key#key IvsPublicKey#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ivs_public_key#value IvsPublicKey#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

