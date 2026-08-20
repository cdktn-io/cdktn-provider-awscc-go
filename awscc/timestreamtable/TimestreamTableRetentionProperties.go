// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreamtable


type TimestreamTableRetentionProperties struct {
	// The duration for which data must be stored in the magnetic store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/timestream_table#magnetic_store_retention_period_in_days TimestreamTable#magnetic_store_retention_period_in_days}
	MagneticStoreRetentionPeriodInDays *string `field:"optional" json:"magneticStoreRetentionPeriodInDays" yaml:"magneticStoreRetentionPeriodInDays"`
	// The duration for which data must be stored in the memory store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/timestream_table#memory_store_retention_period_in_hours TimestreamTable#memory_store_retention_period_in_hours}
	MemoryStoreRetentionPeriodInHours *string `field:"optional" json:"memoryStoreRetentionPeriodInHours" yaml:"memoryStoreRetentionPeriodInHours"`
}

