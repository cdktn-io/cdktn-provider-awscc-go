// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsconfiguredtable

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CleanroomsConfiguredTableConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_configured_table#allowed_columns CleanroomsConfiguredTable#allowed_columns}.
	AllowedColumns *[]*string `field:"required" json:"allowedColumns" yaml:"allowedColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_configured_table#analysis_method CleanroomsConfiguredTable#analysis_method}.
	AnalysisMethod *string `field:"required" json:"analysisMethod" yaml:"analysisMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_configured_table#name CleanroomsConfiguredTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_configured_table#table_reference CleanroomsConfiguredTable#table_reference}.
	TableReference *CleanroomsConfiguredTableTableReference `field:"required" json:"tableReference" yaml:"tableReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_configured_table#analysis_rules CleanroomsConfiguredTable#analysis_rules}.
	AnalysisRules interface{} `field:"optional" json:"analysisRules" yaml:"analysisRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_configured_table#description CleanroomsConfiguredTable#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_configured_table#selected_analysis_methods CleanroomsConfiguredTable#selected_analysis_methods}.
	SelectedAnalysisMethods *[]*string `field:"optional" json:"selectedAnalysisMethods" yaml:"selectedAnalysisMethods"`
	// An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_configured_table#tags CleanroomsConfiguredTable#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

