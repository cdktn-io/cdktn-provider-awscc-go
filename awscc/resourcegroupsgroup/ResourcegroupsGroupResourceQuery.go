// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resourcegroupsgroup


type ResourcegroupsGroupResourceQuery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/resourcegroups_group#query ResourcegroupsGroup#query}.
	Query *ResourcegroupsGroupResourceQueryQuery `field:"optional" json:"query" yaml:"query"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/resourcegroups_group#type ResourcegroupsGroup#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

