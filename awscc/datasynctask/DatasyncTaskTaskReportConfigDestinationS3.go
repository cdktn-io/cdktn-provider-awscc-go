// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynctask


type DatasyncTaskTaskReportConfigDestinationS3 struct {
	// Specifies the Amazon Resource Name (ARN) of the IAM policy that allows Datasync to upload a task report to your S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datasync_task#bucket_access_role_arn DatasyncTask#bucket_access_role_arn}
	BucketAccessRoleArn *string `field:"optional" json:"bucketAccessRoleArn" yaml:"bucketAccessRoleArn"`
	// Specifies the ARN of the S3 bucket where Datasync uploads your report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datasync_task#s3_bucket_arn DatasyncTask#s3_bucket_arn}
	S3BucketArn *string `field:"optional" json:"s3BucketArn" yaml:"s3BucketArn"`
	// Specifies a bucket prefix for your report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datasync_task#subdirectory DatasyncTask#subdirectory}
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
}

