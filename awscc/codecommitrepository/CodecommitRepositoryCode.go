// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codecommitrepository


type CodecommitRepositoryCode struct {
	// Optional.
	//
	// Specifies a branch name to be used as the default branch when importing code into a repository on initial creation. If this property is not set, the name main will be used for the default branch for the repository. Changes to this property are ignored after initial resource creation. We recommend using this parameter to set the name to main to align with the default behavior of CodeCommit unless another name is needed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codecommit_repository#branch_name CodecommitRepository#branch_name}
	BranchName *string `field:"optional" json:"branchName" yaml:"branchName"`
	// Information about the Amazon S3 bucket that contains a ZIP file of code to be committed to the repository.
	//
	// Changes to this property are ignored after initial resource creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codecommit_repository#s3 CodecommitRepository#s3}
	S3 *CodecommitRepositoryCodeS3 `field:"optional" json:"s3" yaml:"s3"`
}

