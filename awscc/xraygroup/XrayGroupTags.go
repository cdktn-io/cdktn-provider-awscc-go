// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package xraygroup


type XrayGroupTags struct {
	// The key name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/xray_group#key XrayGroup#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/xray_group#value XrayGroup#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

