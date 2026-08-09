// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rbinrule


type RbinRuleExcludeResourceTags struct {
	// The tag key of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/rbin_rule#resource_tag_key RbinRule#resource_tag_key}
	ResourceTagKey *string `field:"optional" json:"resourceTagKey" yaml:"resourceTagKey"`
	// The tag value of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/rbin_rule#resource_tag_value RbinRule#resource_tag_value}
	ResourceTagValue *string `field:"optional" json:"resourceTagValue" yaml:"resourceTagValue"`
}

