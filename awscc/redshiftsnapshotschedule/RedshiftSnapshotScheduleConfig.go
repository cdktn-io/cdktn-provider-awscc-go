// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftsnapshotschedule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RedshiftSnapshotScheduleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The definition of the snapshot schedule.
	//
	// The definition is made up of schedule expressions, for example "cron(30 12 *)" or "rate(12 hours)".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/redshift_snapshot_schedule#schedule_definitions RedshiftSnapshotSchedule#schedule_definitions}
	ScheduleDefinitions *[]*string `field:"required" json:"scheduleDefinitions" yaml:"scheduleDefinitions"`
	// A unique identifier for the snapshot schedule. Only alphanumeric characters are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/redshift_snapshot_schedule#schedule_identifier RedshiftSnapshotSchedule#schedule_identifier}
	ScheduleIdentifier *string `field:"required" json:"scheduleIdentifier" yaml:"scheduleIdentifier"`
	// The description of the snapshot schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/redshift_snapshot_schedule#schedule_description RedshiftSnapshotSchedule#schedule_description}
	ScheduleDescription *string `field:"optional" json:"scheduleDescription" yaml:"scheduleDescription"`
	// An optional set of tags for the snapshot schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/redshift_snapshot_schedule#tags RedshiftSnapshotSchedule#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

