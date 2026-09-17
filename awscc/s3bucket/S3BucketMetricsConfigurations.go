// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketMetricsConfigurations struct {
	// The access point that was used while performing operations on the object.
	//
	// The metrics configuration only includes objects that meet the filter's criteria.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#access_point_arn S3Bucket#access_point_arn}
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// The ID used to identify the metrics configuration.
	//
	// This can be any value you choose that helps you identify your metrics configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#id S3Bucket#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The prefix that an object must have to be included in the metrics results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#prefix S3Bucket#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Specifies a list of tag filters to use as a metrics configuration filter.
	//
	// The metrics configuration includes only objects that meet the filter's criteria.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#tag_filters S3Bucket#tag_filters}
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

