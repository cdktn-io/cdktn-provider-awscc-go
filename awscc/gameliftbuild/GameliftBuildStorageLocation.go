// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftbuild


type GameliftBuildStorageLocation struct {
	// An Amazon S3 bucket identifier. This is the name of the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_build#bucket GameliftBuild#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The name of the zip file that contains the build files or script files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_build#key GameliftBuild#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The version of the file, if object versioning is turned on for the bucket.
	//
	// Amazon GameLift uses this information when retrieving files from your S3 bucket. To retrieve a specific version of the file, provide an object version. To retrieve the latest version of the file, do not set this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_build#object_version GameliftBuild#object_version}
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
	// The Amazon Resource Name (ARN) for an IAM role that allows Amazon GameLift to access the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_build#role_arn GameliftBuild#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

