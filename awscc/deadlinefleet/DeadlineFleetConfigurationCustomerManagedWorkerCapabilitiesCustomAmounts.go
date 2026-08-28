// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinefleet


type DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesCustomAmounts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_fleet#max DeadlineFleet#max}.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_fleet#min DeadlineFleet#min}.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_fleet#name DeadlineFleet#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

