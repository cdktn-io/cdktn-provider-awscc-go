// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerAlgorithmConfig struct {
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
	// The name of the algorithm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_algorithm#algorithm_name SagemakerAlgorithm#algorithm_name}
	AlgorithmName *string `field:"required" json:"algorithmName" yaml:"algorithmName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_algorithm#training_specification SagemakerAlgorithm#training_specification}.
	TrainingSpecification *SagemakerAlgorithmTrainingSpecification `field:"required" json:"trainingSpecification" yaml:"trainingSpecification"`
	// A description of the algorithm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_algorithm#algorithm_description SagemakerAlgorithm#algorithm_description}
	AlgorithmDescription *string `field:"optional" json:"algorithmDescription" yaml:"algorithmDescription"`
	// Whether to certify the algorithm so that it can be listed in AWS Marketplace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_algorithm#certify_for_marketplace SagemakerAlgorithm#certify_for_marketplace}
	CertifyForMarketplace interface{} `field:"optional" json:"certifyForMarketplace" yaml:"certifyForMarketplace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_algorithm#inference_specification SagemakerAlgorithm#inference_specification}.
	InferenceSpecification *SagemakerAlgorithmInferenceSpecification `field:"optional" json:"inferenceSpecification" yaml:"inferenceSpecification"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_algorithm#tags SagemakerAlgorithm#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

