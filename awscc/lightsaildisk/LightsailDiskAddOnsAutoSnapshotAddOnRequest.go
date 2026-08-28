// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsaildisk


type LightsailDiskAddOnsAutoSnapshotAddOnRequest struct {
	// The daily time when an automatic snapshot will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_disk#snapshot_time_of_day LightsailDisk#snapshot_time_of_day}
	SnapshotTimeOfDay *string `field:"optional" json:"snapshotTimeOfDay" yaml:"snapshotTimeOfDay"`
}

