// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storagegatewaytape


type StoragegatewayTapeTags struct {
	// The tag key. Cannot be prefixed with aws:.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/storagegateway_tape#key StoragegatewayTape#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/storagegateway_tape#value StoragegatewayTape#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

