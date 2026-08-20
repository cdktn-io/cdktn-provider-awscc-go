// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftscript


type GameliftScriptStorageLocation struct {
	// An Amazon S3 bucket identifier. This is the name of the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_script#bucket GameliftScript#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The name of the zip file that contains the script files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_script#key GameliftScript#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The Amazon Resource Name (ARN) for an IAM role that allows Amazon GameLift to access the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_script#role_arn GameliftScript#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The version of the file, if object versioning is turned on for the bucket.
	//
	// Amazon GameLift uses this information when retrieving files from your S3 bucket. To retrieve a specific version of the file, provide an object version. To retrieve the latest version of the file, do not set this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/gamelift_script#object_version GameliftScript#object_version}
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

