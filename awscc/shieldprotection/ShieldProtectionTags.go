// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package shieldprotection


type ShieldProtectionTags struct {
	// Part of the key:value pair that defines a tag.
	//
	// You can use a tag key to describe a category of information, such as "customer." Tag keys are case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/shield_protection#key ShieldProtection#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Part of the key:value pair that defines a tag.
	//
	// You can use a tag value to describe a specific value within a category, such as "companyA" or "companyB." Tag values are case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/shield_protection#value ShieldProtection#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

