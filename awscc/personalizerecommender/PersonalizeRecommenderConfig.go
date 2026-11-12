// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizerecommender

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PersonalizeRecommenderConfig struct {
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
	// The ARN of the Domain dataset group that the recommender is created for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/personalize_recommender#dataset_group_arn PersonalizeRecommender#dataset_group_arn}
	DatasetGroupArn *string `field:"required" json:"datasetGroupArn" yaml:"datasetGroupArn"`
	// The name of the recommender.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/personalize_recommender#name PersonalizeRecommender#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the recipe (Domain dataset group use case) that the recommender uses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/personalize_recommender#recipe_arn PersonalizeRecommender#recipe_arn}
	RecipeArn *string `field:"required" json:"recipeArn" yaml:"recipeArn"`
	// The configuration details of the recommender.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/personalize_recommender#recommender_config PersonalizeRecommender#recommender_config}
	RecommenderConfig *PersonalizeRecommenderRecommenderConfig `field:"optional" json:"recommenderConfig" yaml:"recommenderConfig"`
	// Tags to associate with the recommender.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/personalize_recommender#tags PersonalizeRecommender#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

