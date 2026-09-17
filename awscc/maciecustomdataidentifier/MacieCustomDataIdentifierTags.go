// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package maciecustomdataidentifier


type MacieCustomDataIdentifierTags struct {
	// The tag's key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/macie_custom_data_identifier#key MacieCustomDataIdentifier#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag's value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/macie_custom_data_identifier#value MacieCustomDataIdentifier#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

