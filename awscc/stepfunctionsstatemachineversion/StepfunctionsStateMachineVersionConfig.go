// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stepfunctionsstatemachineversion

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StepfunctionsStateMachineVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/stepfunctions_state_machine_version#state_machine_arn StepfunctionsStateMachineVersion#state_machine_arn}.
	StateMachineArn *string `field:"required" json:"stateMachineArn" yaml:"stateMachineArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/stepfunctions_state_machine_version#description StepfunctionsStateMachineVersion#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/stepfunctions_state_machine_version#state_machine_revision_id StepfunctionsStateMachineVersion#state_machine_revision_id}.
	StateMachineRevisionId *string `field:"optional" json:"stateMachineRevisionId" yaml:"stateMachineRevisionId"`
}

