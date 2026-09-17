// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendrafeaturedresultsset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KendraFeaturedResultsSetConfig struct {
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
	// A name for the set of featured results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_featured_results_set#featured_results_set_name KendraFeaturedResultsSet#featured_results_set_name}
	FeaturedResultsSetName *string `field:"required" json:"featuredResultsSetName" yaml:"featuredResultsSetName"`
	// The identifier of the index that you want to use for featuring results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_featured_results_set#index_id KendraFeaturedResultsSet#index_id}
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
	// A description for the set of featured results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_featured_results_set#description KendraFeaturedResultsSet#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of document IDs for the documents you want to feature at the top of the search results page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_featured_results_set#featured_documents KendraFeaturedResultsSet#featured_documents}
	FeaturedDocuments interface{} `field:"optional" json:"featuredDocuments" yaml:"featuredDocuments"`
	// A list of queries for featuring results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_featured_results_set#query_texts KendraFeaturedResultsSet#query_texts}
	QueryTexts *[]*string `field:"optional" json:"queryTexts" yaml:"queryTexts"`
	// The current status of the set of featured results.
	//
	// When the value is ACTIVE, featured results are ready for use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_featured_results_set#status KendraFeaturedResultsSet#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// A list of key-value pairs that identify or categorize the featured results set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_featured_results_set#tags KendraFeaturedResultsSet#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

