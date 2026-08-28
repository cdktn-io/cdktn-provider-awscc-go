// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinefleet


type DeadlineFleetConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_fleet#customer_managed DeadlineFleet#customer_managed}.
	CustomerManaged *DeadlineFleetConfigurationCustomerManaged `field:"optional" json:"customerManaged" yaml:"customerManaged"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_fleet#service_managed_ec_2 DeadlineFleet#service_managed_ec_2}.
	ServiceManagedEc2 *DeadlineFleetConfigurationServiceManagedEc2 `field:"optional" json:"serviceManagedEc2" yaml:"serviceManagedEc2"`
}

