// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage


type SagemakerModelPackageModelMetrics struct {
	// Contains bias metrics for a model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_package#bias SagemakerModelPackage#bias}
	Bias *SagemakerModelPackageModelMetricsBias `field:"optional" json:"bias" yaml:"bias"`
	// Contains explainability metrics for a model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_package#explainability SagemakerModelPackage#explainability}
	Explainability *SagemakerModelPackageModelMetricsExplainability `field:"optional" json:"explainability" yaml:"explainability"`
	// Metrics that measure the quality of the input data for a model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_package#model_data_quality SagemakerModelPackage#model_data_quality}
	ModelDataQuality *SagemakerModelPackageModelMetricsModelDataQuality `field:"optional" json:"modelDataQuality" yaml:"modelDataQuality"`
	// Metrics that measure the quality of a model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_package#model_quality SagemakerModelPackage#model_quality}
	ModelQuality *SagemakerModelPackageModelMetricsModelQuality `field:"optional" json:"modelQuality" yaml:"modelQuality"`
}

