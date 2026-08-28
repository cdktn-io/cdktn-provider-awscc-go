// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2ipset


type Wafv2IpSetTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wafv2_ip_set#key Wafv2IpSet#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wafv2_ip_set#value Wafv2IpSet#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

