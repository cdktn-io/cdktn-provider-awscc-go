// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketReplicationConfigurationRulesDestinationReplicationTime struct {
	// Specifies whether the replication time is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3_bucket#status S3Bucket#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// A container specifying the time by which replication should be complete for all objects and operations on objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3_bucket#time S3Bucket#time}
	Time *S3BucketReplicationConfigurationRulesDestinationReplicationTimeTime `field:"optional" json:"time" yaml:"time"`
}

