// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resourcegroupsgroup


type ResourcegroupsGroupResourceQueryQuery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resourcegroups_group#resource_type_filters ResourcegroupsGroup#resource_type_filters}.
	ResourceTypeFilters *[]*string `field:"optional" json:"resourceTypeFilters" yaml:"resourceTypeFilters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resourcegroups_group#stack_identifier ResourcegroupsGroup#stack_identifier}.
	StackIdentifier *string `field:"optional" json:"stackIdentifier" yaml:"stackIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resourcegroups_group#tag_filters ResourcegroupsGroup#tag_filters}.
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

