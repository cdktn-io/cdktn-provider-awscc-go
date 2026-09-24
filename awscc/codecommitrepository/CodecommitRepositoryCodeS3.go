// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codecommitrepository


type CodecommitRepositoryCodeS3 struct {
	// The name of the Amazon S3 bucket that contains the ZIP file with the content that will be committed to the new repository.
	//
	// This can be specified using the name of the bucket in the AWS account. Changes to this property are ignored after initial resource creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codecommit_repository#bucket CodecommitRepository#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The key to use for accessing the Amazon S3 bucket.
	//
	// Changes to this property are ignored after initial resource creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codecommit_repository#key CodecommitRepository#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The object version of the ZIP file, if versioning is enabled for the Amazon S3 bucket.
	//
	// Changes to this property are ignored after initial resource creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codecommit_repository#object_version CodecommitRepository#object_version}
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

