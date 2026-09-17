// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftsnapshotschedule


type RedshiftSnapshotScheduleTags struct {
	// The key, or name, for the resource tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/redshift_snapshot_schedule#key RedshiftSnapshotSchedule#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the resource tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/redshift_snapshot_schedule#value RedshiftSnapshotSchedule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

