// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pricingplanmanagersubscription

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PricingplanmanagerSubscriptionConfig struct {
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
	// The name of the pricing plan family.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/pricingplanmanager_subscription#plan_family PricingplanmanagerSubscription#plan_family}
	PlanFamily *string `field:"required" json:"planFamily" yaml:"planFamily"`
	// The tier of the pricing plan.
	//
	// CloudFormation does not change the tier of an existing subscription; a stack update that changes the tier, upgrading or downgrading it, is rejected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/pricingplanmanager_subscription#plan_tier PricingplanmanagerSubscription#plan_tier}
	PlanTier *string `field:"required" json:"planTier" yaml:"planTier"`
	// The ARNs of resources associated with the subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/pricingplanmanager_subscription#resource_arns PricingplanmanagerSubscription#resource_arns}
	ResourceArns *[]*string `field:"required" json:"resourceArns" yaml:"resourceArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/pricingplanmanager_subscription#usage_level PricingplanmanagerSubscription#usage_level}.
	UsageLevel *string `field:"optional" json:"usageLevel" yaml:"usageLevel"`
}

