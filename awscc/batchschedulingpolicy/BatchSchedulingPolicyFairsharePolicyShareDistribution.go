// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchschedulingpolicy


type BatchSchedulingPolicyFairsharePolicyShareDistribution struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_scheduling_policy#share_identifier BatchSchedulingPolicy#share_identifier}.
	ShareIdentifier *string `field:"optional" json:"shareIdentifier" yaml:"shareIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_scheduling_policy#weight_factor BatchSchedulingPolicy#weight_factor}.
	WeightFactor *float64 `field:"optional" json:"weightFactor" yaml:"weightFactor"`
}

