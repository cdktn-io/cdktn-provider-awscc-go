// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package snstopic


type SnsTopicTags struct {
	// The required key portion of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sns_topic#key SnsTopic#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The optional value portion of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sns_topic#value SnsTopic#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

