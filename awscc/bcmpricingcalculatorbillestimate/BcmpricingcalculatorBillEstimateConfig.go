// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bcmpricingcalculatorbillestimate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BcmpricingcalculatorBillEstimateConfig struct {
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
	// The name of the bill estimate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/bcmpricingcalculator_bill_estimate#name BcmpricingcalculatorBillEstimate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The unique identifier of the bill scenario this estimate is based on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/bcmpricingcalculator_bill_estimate#bill_scenario_id BcmpricingcalculatorBillEstimate#bill_scenario_id}
	BillScenarioId *string `field:"optional" json:"billScenarioId" yaml:"billScenarioId"`
	// The timestamp when the bill estimate expires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/bcmpricingcalculator_bill_estimate#expires_at BcmpricingcalculatorBillEstimate#expires_at}
	ExpiresAt *string `field:"optional" json:"expiresAt" yaml:"expiresAt"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/bcmpricingcalculator_bill_estimate#tags BcmpricingcalculatorBillEstimate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

