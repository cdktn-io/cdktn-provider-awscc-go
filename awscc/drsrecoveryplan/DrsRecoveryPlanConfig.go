// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package drsrecoveryplan

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DrsRecoveryPlanConfig struct {
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
	// The name of the Recovery Plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/drs_recovery_plan#name DrsRecoveryPlan#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the Recovery Plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/drs_recovery_plan#description DrsRecoveryPlan#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags associated with the Recovery Plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/drs_recovery_plan#tags DrsRecoveryPlan#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

