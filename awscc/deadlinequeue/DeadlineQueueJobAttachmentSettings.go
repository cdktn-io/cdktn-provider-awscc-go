// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinequeue


type DeadlineQueueJobAttachmentSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_queue#root_prefix DeadlineQueue#root_prefix}.
	RootPrefix *string `field:"optional" json:"rootPrefix" yaml:"rootPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_queue#s3_bucket_name DeadlineQueue#s3_bucket_name}.
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
}

