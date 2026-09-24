// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsdatamigration


type DmsDataMigrationSourceDataSettings struct {
	// The property is a point in the database engine's log that defines a time where you can begin CDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dms_data_migration#cdc_start_position DmsDataMigration#cdc_start_position}
	CdcStartPosition *string `field:"optional" json:"cdcStartPosition" yaml:"cdcStartPosition"`
	// The property indicates the start time for a change data capture (CDC) operation.
	//
	// The value is server time in UTC format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dms_data_migration#cdc_start_time DmsDataMigration#cdc_start_time}
	CdcStartTime *string `field:"optional" json:"cdcStartTime" yaml:"cdcStartTime"`
	// The property indicates the stop time for a change data capture (CDC) operation.
	//
	// The value is server time in UTC format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dms_data_migration#cdc_stop_time DmsDataMigration#cdc_stop_time}
	CdcStopTime *string `field:"optional" json:"cdcStopTime" yaml:"cdcStopTime"`
	// The property sets the name of a previously created logical replication slot for a change data capture (CDC) load of the source instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dms_data_migration#slot_name DmsDataMigration#slot_name}
	SlotName *string `field:"optional" json:"slotName" yaml:"slotName"`
}

