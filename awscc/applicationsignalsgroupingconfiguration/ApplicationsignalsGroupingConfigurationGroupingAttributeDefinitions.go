// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsgroupingconfiguration


type ApplicationsignalsGroupingConfigurationGroupingAttributeDefinitions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/applicationsignals_grouping_configuration#grouping_name ApplicationsignalsGroupingConfiguration#grouping_name}.
	GroupingName *string `field:"required" json:"groupingName" yaml:"groupingName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/applicationsignals_grouping_configuration#grouping_source_keys ApplicationsignalsGroupingConfiguration#grouping_source_keys}.
	GroupingSourceKeys *[]*string `field:"required" json:"groupingSourceKeys" yaml:"groupingSourceKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/applicationsignals_grouping_configuration#default_grouping_value ApplicationsignalsGroupingConfiguration#default_grouping_value}.
	DefaultGroupingValue *string `field:"optional" json:"defaultGroupingValue" yaml:"defaultGroupingValue"`
}

