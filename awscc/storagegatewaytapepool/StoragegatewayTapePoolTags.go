// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storagegatewaytapepool


type StoragegatewayTapePoolTags struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/storagegateway_tape_pool#key StoragegatewayTapePool#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/storagegateway_tape_pool#value StoragegatewayTapePool#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

