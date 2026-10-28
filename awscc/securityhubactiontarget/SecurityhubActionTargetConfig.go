// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubactiontarget

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityhubActionTargetConfig struct {
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
	// The ID for the custom action target. Can contain up to 20 alphanumeric characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityhub_action_target#action_target_id SecurityhubActionTarget#action_target_id}
	ActionTargetId *string `field:"required" json:"actionTargetId" yaml:"actionTargetId"`
	// The description for the custom action target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityhub_action_target#description SecurityhubActionTarget#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the custom action target. Can contain up to 20 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityhub_action_target#name SecurityhubActionTarget#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

