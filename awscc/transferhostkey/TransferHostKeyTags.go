// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferhostkey


type TransferHostKeyTags struct {
	// The name assigned to the tag that you create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/transfer_host_key#key TransferHostKey#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Contains one or more values that you assigned to the key name you create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/transfer_host_key#value TransferHostKey#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

