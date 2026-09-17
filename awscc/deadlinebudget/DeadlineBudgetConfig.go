// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinebudget

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DeadlineBudgetConfig struct {
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
	// The budget actions to specify what happens when the budget runs out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_budget#actions DeadlineBudget#actions}
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The dollar limit based on consumed usage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_budget#approximate_dollar_limit DeadlineBudget#approximate_dollar_limit}
	ApproximateDollarLimit *float64 `field:"required" json:"approximateDollarLimit" yaml:"approximateDollarLimit"`
	// The display name of the budget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_budget#display_name DeadlineBudget#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The farm ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_budget#farm_id DeadlineBudget#farm_id}
	FarmId *string `field:"required" json:"farmId" yaml:"farmId"`
	// The start and end time of the budget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_budget#schedule DeadlineBudget#schedule}
	Schedule *DeadlineBudgetSchedule `field:"required" json:"schedule" yaml:"schedule"`
	// The usage details of the allotted budget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_budget#usage_tracking_resource DeadlineBudget#usage_tracking_resource}
	UsageTrackingResource *DeadlineBudgetUsageTrackingResource `field:"required" json:"usageTrackingResource" yaml:"usageTrackingResource"`
	// The description of the budget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_budget#description DeadlineBudget#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_budget#tags DeadlineBudget#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

