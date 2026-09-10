// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonedatasource


type DatazoneDataSourceSchedule struct {
	// The schedule of the data source runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_data_source#schedule DatazoneDataSource#schedule}
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// The timezone of the data source run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_data_source#timezone DatazoneDataSource#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

