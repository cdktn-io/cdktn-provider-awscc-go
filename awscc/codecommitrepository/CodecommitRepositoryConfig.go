// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codecommitrepository

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CodecommitRepositoryConfig struct {
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
	// The name of the new repository to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codecommit_repository#repository_name CodecommitRepository#repository_name}
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// Information about code to be committed to a repository after it is created in an AWS CloudFormation stack.
	//
	// Information about code is only used in resource creation. Updates to a stack will not reflect changes made to code properties after initial resource creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codecommit_repository#code CodecommitRepository#code}
	Code *CodecommitRepositoryCode `field:"optional" json:"code" yaml:"code"`
	// The ID of the AWS Key Management Service encryption key used to encrypt and decrypt the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codecommit_repository#kms_key_id CodecommitRepository#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// A comment or description about the new repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codecommit_repository#repository_description CodecommitRepository#repository_description}
	RepositoryDescription *string `field:"optional" json:"repositoryDescription" yaml:"repositoryDescription"`
	// One or more tag key-value pairs to use when tagging this repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codecommit_repository#tags CodecommitRepository#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Information about a trigger for a repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codecommit_repository#triggers CodecommitRepository#triggers}
	Triggers interface{} `field:"optional" json:"triggers" yaml:"triggers"`
}

