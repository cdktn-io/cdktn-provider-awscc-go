// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bcmpricingcalculatorbillscenario

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BcmpricingcalculatorBillScenarioConfig struct {
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
	// The ARN of the cost category group sharing preference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bcmpricingcalculator_bill_scenario#cost_category_group_sharing_preference_arn BcmpricingcalculatorBillScenario#cost_category_group_sharing_preference_arn}
	CostCategoryGroupSharingPreferenceArn *string `field:"optional" json:"costCategoryGroupSharingPreferenceArn" yaml:"costCategoryGroupSharingPreferenceArn"`
	// The timestamp when the bill scenario expires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bcmpricingcalculator_bill_scenario#expires_at BcmpricingcalculatorBillScenario#expires_at}
	ExpiresAt *string `field:"optional" json:"expiresAt" yaml:"expiresAt"`
	// The group sharing preference for the bill scenario.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bcmpricingcalculator_bill_scenario#group_sharing_preference BcmpricingcalculatorBillScenario#group_sharing_preference}
	GroupSharingPreference *string `field:"optional" json:"groupSharingPreference" yaml:"groupSharingPreference"`
	// The name of the bill scenario.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bcmpricingcalculator_bill_scenario#name BcmpricingcalculatorBillScenario#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bcmpricingcalculator_bill_scenario#tags BcmpricingcalculatorBillScenario#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

