// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkflowmonitorscope


type NetworkflowmonitorScopeTargetsTargetIdentifier struct {
	// A target ID is an internally-generated identifier for a target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/networkflowmonitor_scope#target_id NetworkflowmonitorScope#target_id}
	TargetId *NetworkflowmonitorScopeTargetsTargetIdentifierTargetId `field:"required" json:"targetId" yaml:"targetId"`
	// The type of the target. Currently always ACCOUNT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/networkflowmonitor_scope#target_type NetworkflowmonitorScope#target_type}
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
}

