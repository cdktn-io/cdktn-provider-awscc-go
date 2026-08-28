// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2service


type Resiliencehubv2ServiceReportConfigurationReportOutputS3 struct {
	// Account ID of the bucket owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/resiliencehubv2_service#bucket_owner Resiliencehubv2Service#bucket_owner}
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// S3 bucket path where reports will be written.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/resiliencehubv2_service#bucket_path Resiliencehubv2Service#bucket_path}
	BucketPath *string `field:"optional" json:"bucketPath" yaml:"bucketPath"`
}

