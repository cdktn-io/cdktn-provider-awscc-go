// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkflowmonitormonitor


type NetworkflowmonitorMonitorRemoteResources struct {
	// The identifier of the remote resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/networkflowmonitor_monitor#identifier NetworkflowmonitorMonitor#identifier}
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// The type of the remote resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/networkflowmonitor_monitor#type NetworkflowmonitorMonitor#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

