// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchservicedomain


type OpensearchserviceDomainSoftwareUpdateOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchservice_domain#auto_software_update_enabled OpensearchserviceDomain#auto_software_update_enabled}.
	AutoSoftwareUpdateEnabled interface{} `field:"optional" json:"autoSoftwareUpdateEnabled" yaml:"autoSoftwareUpdateEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchservice_domain#use_latest_service_software_for_blue_green OpensearchserviceDomain#use_latest_service_software_for_blue_green}.
	UseLatestServiceSoftwareForBlueGreen interface{} `field:"optional" json:"useLatestServiceSoftwareForBlueGreen" yaml:"useLatestServiceSoftwareForBlueGreen"`
}

