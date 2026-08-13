// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynctask


type DatasyncTaskManifestConfigSourceS3 struct {
	// Specifies the AWS Identity and Access Management (IAM) role that allows DataSync to access your manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datasync_task#bucket_access_role_arn DatasyncTask#bucket_access_role_arn}
	BucketAccessRoleArn *string `field:"optional" json:"bucketAccessRoleArn" yaml:"bucketAccessRoleArn"`
	// Specifies the Amazon S3 object key of your manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datasync_task#manifest_object_path DatasyncTask#manifest_object_path}
	ManifestObjectPath *string `field:"optional" json:"manifestObjectPath" yaml:"manifestObjectPath"`
	// Specifies the object version ID of the manifest that you want DataSync to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datasync_task#manifest_object_version_id DatasyncTask#manifest_object_version_id}
	ManifestObjectVersionId *string `field:"optional" json:"manifestObjectVersionId" yaml:"manifestObjectVersionId"`
	// Specifies the Amazon Resource Name (ARN) of the S3 bucket where you're hosting your manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datasync_task#s3_bucket_arn DatasyncTask#s3_bucket_arn}
	S3BucketArn *string `field:"optional" json:"s3BucketArn" yaml:"s3BucketArn"`
}

