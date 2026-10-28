// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingplansscalingplan

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AutoscalingplansScalingPlanConfig struct {
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
	// A CloudFormation stack or a set of tags.
	//
	// You can create one scaling plan per application source. The ApplicationSource property must be present to ensure interoperability with the AWS Auto Scaling console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#application_source AutoscalingplansScalingPlan#application_source}
	ApplicationSource *AutoscalingplansScalingPlanApplicationSource `field:"required" json:"applicationSource" yaml:"applicationSource"`
	// The scaling instructions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#scaling_instructions AutoscalingplansScalingPlan#scaling_instructions}
	ScalingInstructions interface{} `field:"required" json:"scalingInstructions" yaml:"scalingInstructions"`
}

