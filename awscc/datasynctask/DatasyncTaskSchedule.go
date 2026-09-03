// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynctask


type DatasyncTaskSchedule struct {
	// A cron expression that specifies when AWS DataSync initiates a scheduled transfer from a source to a destination location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datasync_task#schedule_expression DatasyncTask#schedule_expression}
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
	// Specifies status of a schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datasync_task#status DatasyncTask#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

