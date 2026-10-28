// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercomputequota


type SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfig struct {
	// Per-instance-type absolute caps on resources this team can borrow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#absolute_borrow_limits SagemakerComputeQuota#absolute_borrow_limits}
	AbsoluteBorrowLimits interface{} `field:"optional" json:"absoluteBorrowLimits" yaml:"absoluteBorrowLimits"`
	// Maximum percentage of the team's idle quota that can be borrowed by other teams.
	//
	// Smithy admits 0..10000; in practice the service-documented operational range is 1..500. Defaults service-side to 50 when omitted on Create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#borrow_limit SagemakerComputeQuota#borrow_limit}
	BorrowLimit *float64 `field:"optional" json:"borrowLimit" yaml:"borrowLimit"`
	// How this team participates in idle-resource sharing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#strategy SagemakerComputeQuota#strategy}
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

