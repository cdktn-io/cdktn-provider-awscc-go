// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendraquerysuggestionsblocklist

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KendraQuerySuggestionsBlockListConfig struct {
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
	// The identifier of the index for the block list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kendra_query_suggestions_block_list#index_id KendraQuerySuggestionsBlockList#index_id}
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
	// The name of the block list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kendra_query_suggestions_block_list#name KendraQuerySuggestionsBlockList#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of an IAM role with permission to access the S3 bucket that contains the block list text file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kendra_query_suggestions_block_list#role_arn KendraQuerySuggestionsBlockList#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Information required to find a specific file in an Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kendra_query_suggestions_block_list#source_s3_path KendraQuerySuggestionsBlockList#source_s3_path}
	SourceS3Path *KendraQuerySuggestionsBlockListSourceS3Path `field:"required" json:"sourceS3Path" yaml:"sourceS3Path"`
	// A description for the block list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kendra_query_suggestions_block_list#description KendraQuerySuggestionsBlockList#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of key-value pairs that identify or categorize the block list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kendra_query_suggestions_block_list#tags KendraQuerySuggestionsBlockList#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

