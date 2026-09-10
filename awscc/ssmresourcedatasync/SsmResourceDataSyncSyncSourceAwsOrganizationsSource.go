// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmresourcedatasync


type SsmResourceDataSyncSyncSourceAwsOrganizationsSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssm_resource_data_sync#organizational_units SsmResourceDataSync#organizational_units}.
	OrganizationalUnits *[]*string `field:"optional" json:"organizationalUnits" yaml:"organizationalUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssm_resource_data_sync#organization_source_type SsmResourceDataSync#organization_source_type}.
	OrganizationSourceType *string `field:"optional" json:"organizationSourceType" yaml:"organizationSourceType"`
}

