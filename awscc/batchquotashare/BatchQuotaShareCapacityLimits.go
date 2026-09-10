// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchquotashare


type BatchQuotaShareCapacityLimits struct {
	// The unit of compute capacity for the capacityLimit. For example, `ml.m5.large`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_quota_share#capacity_unit BatchQuotaShare#capacity_unit}
	CapacityUnit *string `field:"required" json:"capacityUnit" yaml:"capacityUnit"`
	// The maximum capacity available for the quota share.
	//
	// This value represents the maximum quantity of a resource that can be allocated to jobs in the quota share without borrowing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_quota_share#max_capacity BatchQuotaShare#max_capacity}
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
}

