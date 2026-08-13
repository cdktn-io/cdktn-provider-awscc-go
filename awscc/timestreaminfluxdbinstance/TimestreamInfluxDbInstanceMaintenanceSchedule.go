// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreaminfluxdbinstance


type TimestreamInfluxDbInstanceMaintenanceSchedule struct {
	// The preferred maintenance window in format ddd:HH:MM-ddd:HH:MM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#preferred_maintenance_window TimestreamInfluxDbInstance#preferred_maintenance_window}
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The IANA timezone identifier for the maintenance schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#timezone TimestreamInfluxDbInstance#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

