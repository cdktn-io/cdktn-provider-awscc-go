// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimeappinstance


type ChimeAppInstanceTags struct {
	// The key in a tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/chime_app_instance#key ChimeAppInstance#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value in a tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/chime_app_instance#value ChimeAppInstance#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

