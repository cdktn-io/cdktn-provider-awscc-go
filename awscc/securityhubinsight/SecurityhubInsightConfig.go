// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubinsight

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityhubInsightConfig struct {
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
	// One or more attributes used to filter the findings included in the insight.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/securityhub_insight#filters SecurityhubInsight#filters}
	Filters *SecurityhubInsightFilters `field:"required" json:"filters" yaml:"filters"`
	// The grouping attribute for the insight's findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/securityhub_insight#group_by_attribute SecurityhubInsight#group_by_attribute}
	GroupByAttribute *string `field:"required" json:"groupByAttribute" yaml:"groupByAttribute"`
	// The name of a Security Hub insight.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/securityhub_insight#name SecurityhubInsight#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

