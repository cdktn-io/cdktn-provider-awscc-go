// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinefleet


type DeadlineFleetConfigurationServiceManagedEc2VpcConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_fleet#resource_configuration_arns DeadlineFleet#resource_configuration_arns}.
	ResourceConfigurationArns *[]*string `field:"optional" json:"resourceConfigurationArns" yaml:"resourceConfigurationArns"`
}

