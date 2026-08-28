// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package controltowerenabledcontrol

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ControltowerEnabledControlConfig struct {
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
	// Arn of the control.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/controltower_enabled_control#control_identifier ControltowerEnabledControl#control_identifier}
	ControlIdentifier *string `field:"required" json:"controlIdentifier" yaml:"controlIdentifier"`
	// Arn for Organizational unit to which the control needs to be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/controltower_enabled_control#target_identifier ControltowerEnabledControl#target_identifier}
	TargetIdentifier *string `field:"required" json:"targetIdentifier" yaml:"targetIdentifier"`
}

