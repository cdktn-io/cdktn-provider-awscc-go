// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage


type SagemakerModelPackageSourceAlgorithmSpecificationSourceAlgorithms struct {
	// The name of an algorithm that was used to create the model package.
	//
	// The algorithm must be either an algorithm resource in your Amazon SageMaker account or an algorithm in AWS Marketplace that you are subscribed to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_model_package#algorithm_name SagemakerModelPackage#algorithm_name}
	AlgorithmName *string `field:"optional" json:"algorithmName" yaml:"algorithmName"`
	// The Amazon S3 path where the model artifacts, which result from model training, are stored.
	//
	// This path must point to a single gzip compressed tar archive (.tar.gz suffix).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_model_package#model_data_url SagemakerModelPackage#model_data_url}
	ModelDataUrl *string `field:"optional" json:"modelDataUrl" yaml:"modelDataUrl"`
}

