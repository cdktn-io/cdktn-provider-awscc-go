// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmresourcedatasync


type SsmResourceDataSyncSyncSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_resource_data_sync#aws_organizations_source SsmResourceDataSync#aws_organizations_source}.
	AwsOrganizationsSource *SsmResourceDataSyncSyncSourceAwsOrganizationsSource `field:"optional" json:"awsOrganizationsSource" yaml:"awsOrganizationsSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_resource_data_sync#include_future_regions SsmResourceDataSync#include_future_regions}.
	IncludeFutureRegions interface{} `field:"optional" json:"includeFutureRegions" yaml:"includeFutureRegions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_resource_data_sync#source_regions SsmResourceDataSync#source_regions}.
	SourceRegions *[]*string `field:"optional" json:"sourceRegions" yaml:"sourceRegions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_resource_data_sync#source_type SsmResourceDataSync#source_type}.
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}

