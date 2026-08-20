// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilesrecommender

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CustomerprofilesRecommenderConfig struct {
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
	// The name of the domain for which the recommender will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_recommender#domain_name CustomerprofilesRecommender#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The name of the recommender.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_recommender#recommender_name CustomerprofilesRecommender#recommender_name}
	RecommenderName *string `field:"required" json:"recommenderName" yaml:"recommenderName"`
	// The name of the recommender recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_recommender#recommender_recipe_name CustomerprofilesRecommender#recommender_recipe_name}
	RecommenderRecipeName *string `field:"required" json:"recommenderRecipeName" yaml:"recommenderRecipeName"`
	// The description of the recommender.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_recommender#description CustomerprofilesRecommender#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configuration for the recommender.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_recommender#recommender_config CustomerprofilesRecommender#recommender_config}
	RecommenderConfig *CustomerprofilesRecommenderRecommenderConfig `field:"optional" json:"recommenderConfig" yaml:"recommenderConfig"`
	// The tags used to organize, track, or control access for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_recommender#tags CustomerprofilesRecommender#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

