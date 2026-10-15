// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilesrecommenderfilter

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CustomerprofilesRecommenderFilterConfig struct {
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
	// The unique name of the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/customerprofiles_recommender_filter#domain_name CustomerprofilesRecommenderFilter#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The filter expression that defines which items to include or exclude from recommendations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/customerprofiles_recommender_filter#recommender_filter_expression CustomerprofilesRecommenderFilter#recommender_filter_expression}
	RecommenderFilterExpression *string `field:"required" json:"recommenderFilterExpression" yaml:"recommenderFilterExpression"`
	// The name of the recommender filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/customerprofiles_recommender_filter#recommender_filter_name CustomerprofilesRecommenderFilter#recommender_filter_name}
	RecommenderFilterName *string `field:"required" json:"recommenderFilterName" yaml:"recommenderFilterName"`
	// The description of the recommender filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/customerprofiles_recommender_filter#description CustomerprofilesRecommenderFilter#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the recommender schema to use for this recommender filter.
	//
	// If not specified, the default schema is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/customerprofiles_recommender_filter#recommender_schema_name CustomerprofilesRecommenderFilter#recommender_schema_name}
	RecommenderSchemaName *string `field:"optional" json:"recommenderSchemaName" yaml:"recommenderSchemaName"`
	// The tags used to organize, track, or control access for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/customerprofiles_recommender_filter#tags CustomerprofilesRecommenderFilter#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

