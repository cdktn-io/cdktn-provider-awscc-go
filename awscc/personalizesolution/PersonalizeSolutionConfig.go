// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizesolution

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PersonalizeSolutionConfig struct {
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
	// The ARN of the dataset group that provides the training data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/personalize_solution#dataset_group_arn PersonalizeSolution#dataset_group_arn}
	DatasetGroupArn *string `field:"required" json:"datasetGroupArn" yaml:"datasetGroupArn"`
	// The name for the solution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/personalize_solution#name PersonalizeSolution#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// When your have multiple event types (using an EVENT_TYPE schema field), this parameter specifies which event type (for example, 'click' or 'like') is used for training the model.
	//
	// If you do not provide an eventType, Amazon Personalize will use all interactions for training with equal weight regardless of type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/personalize_solution#event_type PersonalizeSolution#event_type}
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// Whether to perform automated machine learning (AutoML). The default is false. For this case, you must specify recipeArn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/personalize_solution#perform_auto_ml PersonalizeSolution#perform_auto_ml}
	PerformAutoMl interface{} `field:"optional" json:"performAutoMl" yaml:"performAutoMl"`
	// Whether to perform hyperparameter optimization (HPO) on the specified or selected recipe.
	//
	// The default is false. When performing AutoML, this parameter is always true and you should not set it to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/personalize_solution#perform_hpo PersonalizeSolution#perform_hpo}
	PerformHpo interface{} `field:"optional" json:"performHpo" yaml:"performHpo"`
	// The ARN of the recipe to use for model training. Only specified when performAutoML is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/personalize_solution#recipe_arn PersonalizeSolution#recipe_arn}
	RecipeArn *string `field:"optional" json:"recipeArn" yaml:"recipeArn"`
	// The configuration to use with the solution.
	//
	// When performAutoML is set to true, Amazon Personalize only evaluates the autoMLConfig section of the solution configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/personalize_solution#solution_config PersonalizeSolution#solution_config}
	SolutionConfig *PersonalizeSolutionSolutionConfig `field:"optional" json:"solutionConfig" yaml:"solutionConfig"`
	// The tags used to organize, track, or control access for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/personalize_solution#tags PersonalizeSolution#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

