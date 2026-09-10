// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3expressaccesspoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3ExpressAccessPointConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the bucket that you want to associate this Access Point with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/s3express_access_point#bucket S3ExpressAccessPoint#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The AWS account ID associated with the S3 bucket associated with this access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/s3express_access_point#bucket_account_id S3ExpressAccessPoint#bucket_account_id}
	BucketAccountId *string `field:"optional" json:"bucketAccountId" yaml:"bucketAccountId"`
	// The name you want to assign to this Access Point.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name. For directory buckets, the access point name must consist of a base name that you provide and su?x that includes the ZoneID (AWS Availability Zone or Local Zone) of your bucket location, followed by --xa-s3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/s3express_access_point#name S3ExpressAccessPoint#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Access Point Policy you want to apply to this access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/s3express_access_point#policy S3ExpressAccessPoint#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// The PublicAccessBlock configuration that you want to apply to this Access Point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/s3express_access_point#public_access_block_configuration S3ExpressAccessPoint#public_access_block_configuration}
	PublicAccessBlockConfiguration *S3ExpressAccessPointPublicAccessBlockConfiguration `field:"optional" json:"publicAccessBlockConfiguration" yaml:"publicAccessBlockConfiguration"`
	// For directory buckets, you can ?lter access control to speci?c pre?xes, API operations, or a combination of both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/s3express_access_point#scope S3ExpressAccessPoint#scope}
	Scope *S3ExpressAccessPointScope `field:"optional" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/s3express_access_point#tags S3ExpressAccessPoint#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// If you include this field, Amazon S3 restricts access to this Access Point to requests from the specified Virtual Private Cloud (VPC).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/s3express_access_point#vpc_configuration S3ExpressAccessPoint#vpc_configuration}
	VpcConfiguration *S3ExpressAccessPointVpcConfiguration `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

