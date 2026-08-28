// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2verifiedaccessinstance


type Ec2VerifiedAccessInstanceLoggingConfigurationsS3 struct {
	// The bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_verified_access_instance#bucket_name Ec2VerifiedAccessInstance#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The ID of the AWS account that owns the Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_verified_access_instance#bucket_owner Ec2VerifiedAccessInstance#bucket_owner}
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// Indicates whether logging is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_verified_access_instance#enabled Ec2VerifiedAccessInstance#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The bucket prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_verified_access_instance#prefix Ec2VerifiedAccessInstance#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

