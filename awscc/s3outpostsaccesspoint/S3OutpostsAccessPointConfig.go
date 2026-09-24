// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3outpostsaccesspoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3OutpostsAccessPointConfig struct {
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
	// The Amazon Resource Name (ARN) of the bucket you want to associate this AccessPoint with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3outposts_access_point#bucket S3OutpostsAccessPoint#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// A name for the AccessPoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3outposts_access_point#name S3OutpostsAccessPoint#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Virtual Private Cloud (VPC) from which requests can be made to the AccessPoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3outposts_access_point#vpc_configuration S3OutpostsAccessPoint#vpc_configuration}
	VpcConfiguration *S3OutpostsAccessPointVpcConfiguration `field:"required" json:"vpcConfiguration" yaml:"vpcConfiguration"`
	// The access point policy associated with this access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3outposts_access_point#policy S3OutpostsAccessPoint#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

