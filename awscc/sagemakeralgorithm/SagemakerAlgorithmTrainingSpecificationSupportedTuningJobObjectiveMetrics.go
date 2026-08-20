// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm


type SagemakerAlgorithmTrainingSpecificationSupportedTuningJobObjectiveMetrics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_algorithm#metric_name SagemakerAlgorithm#metric_name}.
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_algorithm#type SagemakerAlgorithm#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

