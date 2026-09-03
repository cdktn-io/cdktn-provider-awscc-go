// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluedataqualityruleset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueDataQualityRulesetConfig struct {
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
	// A unique name for the data quality ruleset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_data_quality_ruleset#name GlueDataQualityRuleset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A unique token for idempotency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_data_quality_ruleset#client_token GlueDataQualityRuleset#client_token}
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// A description of the data quality ruleset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_data_quality_ruleset#description GlueDataQualityRuleset#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A Data Quality Definition Language (DQDL) ruleset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_data_quality_ruleset#ruleset GlueDataQualityRuleset#ruleset}
	Ruleset *string `field:"optional" json:"ruleset" yaml:"ruleset"`
	// A map of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_data_quality_ruleset#tags GlueDataQualityRuleset#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// An object representing an AWS Glue table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_data_quality_ruleset#target_table GlueDataQualityRuleset#target_table}
	TargetTable *GlueDataQualityRulesetTargetTable `field:"optional" json:"targetTable" yaml:"targetTable"`
}

