// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkflowmonitorscope


type NetworkflowmonitorScopeTargets struct {
	// The AWS Region for the target resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/networkflowmonitor_scope#region NetworkflowmonitorScope#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// A target identifier is a pair of identifying information for a scope target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/networkflowmonitor_scope#target_identifier NetworkflowmonitorScope#target_identifier}
	TargetIdentifier *NetworkflowmonitorScopeTargetsTargetIdentifier `field:"required" json:"targetIdentifier" yaml:"targetIdentifier"`
}

