// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm


type SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRangeIntegerParameterRangeSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_algorithm#max_value SagemakerAlgorithm#max_value}.
	MaxValue *string `field:"optional" json:"maxValue" yaml:"maxValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_algorithm#min_value SagemakerAlgorithm#min_value}.
	MinValue *string `field:"optional" json:"minValue" yaml:"minValue"`
}

