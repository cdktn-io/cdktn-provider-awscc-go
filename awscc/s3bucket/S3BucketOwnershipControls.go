// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketOwnershipControls struct {
	// Specifies the container element for Object Ownership rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#rules S3Bucket#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

