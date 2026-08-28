// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkflowmonitormonitor


type NetworkflowmonitorMonitorLocalResources struct {
	// The identifier of the local resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkflowmonitor_monitor#identifier NetworkflowmonitorMonitor#identifier}
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// The type of the local resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkflowmonitor_monitor#type NetworkflowmonitorMonitor#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

