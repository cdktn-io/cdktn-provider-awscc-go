// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection


type DatazoneConnectionPropsRedshiftPropertiesLineageSync struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#enabled DatazoneConnection#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Lineage Sync Schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#schedule DatazoneConnection#schedule}
	Schedule *DatazoneConnectionPropsRedshiftPropertiesLineageSyncSchedule `field:"optional" json:"schedule" yaml:"schedule"`
}

