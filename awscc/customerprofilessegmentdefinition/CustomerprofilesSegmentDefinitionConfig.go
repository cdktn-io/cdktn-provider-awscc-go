// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilessegmentdefinition

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CustomerprofilesSegmentDefinitionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The display name of the segment definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_segment_definition#display_name CustomerprofilesSegmentDefinition#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The unique name of the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_segment_definition#domain_name CustomerprofilesSegmentDefinition#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The unique name of the segment definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_segment_definition#segment_definition_name CustomerprofilesSegmentDefinition#segment_definition_name}
	SegmentDefinitionName *string `field:"required" json:"segmentDefinitionName" yaml:"segmentDefinitionName"`
	// The description of the segment definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_segment_definition#description CustomerprofilesSegmentDefinition#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array that defines the set of segment criteria to evaluate when handling segment groups for the segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_segment_definition#segment_groups CustomerprofilesSegmentDefinition#segment_groups}
	SegmentGroups *CustomerprofilesSegmentDefinitionSegmentGroups `field:"optional" json:"segmentGroups" yaml:"segmentGroups"`
	// The segment sort configuration for ordering segment results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_segment_definition#segment_sort CustomerprofilesSegmentDefinition#segment_sort}
	SegmentSort *CustomerprofilesSegmentDefinitionSegmentSort `field:"optional" json:"segmentSort" yaml:"segmentSort"`
	// The SQL query that defines the segment criteria.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_segment_definition#segment_sql_query CustomerprofilesSegmentDefinition#segment_sql_query}
	SegmentSqlQuery *string `field:"optional" json:"segmentSqlQuery" yaml:"segmentSqlQuery"`
	// The tags used to organize, track, or control access for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_segment_definition#tags CustomerprofilesSegmentDefinition#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

