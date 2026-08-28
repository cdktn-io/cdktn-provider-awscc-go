// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferagreement


type TransferAgreementTags struct {
	// The name assigned to the tag that you create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/transfer_agreement#key TransferAgreement#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Contains one or more values that you assigned to the key name you create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/transfer_agreement#value TransferAgreement#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

