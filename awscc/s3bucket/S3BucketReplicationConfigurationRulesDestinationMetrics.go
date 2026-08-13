// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketReplicationConfigurationRulesDestinationMetrics struct {
	// A container specifying the time threshold for emitting the ``s3:Replication:OperationMissedThreshold`` event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3_bucket#event_threshold S3Bucket#event_threshold}
	EventThreshold *S3BucketReplicationConfigurationRulesDestinationMetricsEventThreshold `field:"optional" json:"eventThreshold" yaml:"eventThreshold"`
	// Specifies whether the replication metrics are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3_bucket#status S3Bucket#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

