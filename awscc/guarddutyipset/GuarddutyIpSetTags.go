// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package guarddutyipset


type GuarddutyIpSetTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/guardduty_ip_set#key GuarddutyIpSet#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/guardduty_ip_set#value GuarddutyIpSet#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

