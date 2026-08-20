// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchquotashare


type BatchQuotaSharePreemptionConfiguration struct {
	// Specifies whether jobs within a quota share can be preempted by another, higher priority job in the same quota share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_quota_share#in_share_preemption BatchQuotaShare#in_share_preemption}
	InSharePreemption *string `field:"required" json:"inSharePreemption" yaml:"inSharePreemption"`
}

