// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinefleet


type DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesAcceleratorCapabilities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_fleet#count DeadlineFleet#count}.
	Count *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesAcceleratorCapabilitiesCount `field:"optional" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_fleet#selections DeadlineFleet#selections}.
	Selections interface{} `field:"optional" json:"selections" yaml:"selections"`
}

