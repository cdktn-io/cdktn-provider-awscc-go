// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3expressdirectorybucket


type S3ExpressDirectoryBucketMetricsConfigurations struct {
	// The access point ARN used when evaluating a metrics filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3express_directory_bucket#access_point_arn S3ExpressDirectoryBucket#access_point_arn}
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// The ID used to identify the metrics configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3express_directory_bucket#id S3ExpressDirectoryBucket#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The prefix used when evaluating a metrics filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3express_directory_bucket#prefix S3ExpressDirectoryBucket#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

