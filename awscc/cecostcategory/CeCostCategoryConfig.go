// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cecostcategory

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CeCostCategoryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ce_cost_category#name CeCostCategory#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// JSON array format of Expression in Billing and Cost Management API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ce_cost_category#rules CeCostCategory#rules}
	Rules *string `field:"required" json:"rules" yaml:"rules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ce_cost_category#rule_version CeCostCategory#rule_version}.
	RuleVersion *string `field:"required" json:"ruleVersion" yaml:"ruleVersion"`
	// The default value for the cost category.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ce_cost_category#default_value CeCostCategory#default_value}
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// Json array format of CostCategorySplitChargeRule in Billing and Cost Management API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ce_cost_category#split_charge_rules CeCostCategory#split_charge_rules}
	SplitChargeRules *string `field:"optional" json:"splitChargeRules" yaml:"splitChargeRules"`
	// Tags to assign to the cost category.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ce_cost_category#tags CeCostCategory#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

