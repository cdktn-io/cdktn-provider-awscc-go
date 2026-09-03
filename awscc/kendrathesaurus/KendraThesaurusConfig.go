// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendrathesaurus

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KendraThesaurusConfig struct {
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
	// The identifier of the index for the thesaurus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_thesaurus#index_id KendraThesaurus#index_id}
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
	// A name for the thesaurus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_thesaurus#name KendraThesaurus#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An IAM role that gives Amazon Kendra permissions to access the thesaurus file specified in SourceS3Path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_thesaurus#role_arn KendraThesaurus#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Information required to find a specific file in an Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_thesaurus#source_s3_path KendraThesaurus#source_s3_path}
	SourceS3Path *KendraThesaurusSourceS3Path `field:"required" json:"sourceS3Path" yaml:"sourceS3Path"`
	// A description for the thesaurus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_thesaurus#description KendraThesaurus#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of key-value pairs that identify or categorize the thesaurus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_thesaurus#tags KendraThesaurus#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

