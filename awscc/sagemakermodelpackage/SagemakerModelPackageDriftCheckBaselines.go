// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage


type SagemakerModelPackageDriftCheckBaselines struct {
	// Represents the drift check bias baselines that can be used when the model monitor is set using the model package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_model_package#bias SagemakerModelPackage#bias}
	Bias *SagemakerModelPackageDriftCheckBaselinesBias `field:"optional" json:"bias" yaml:"bias"`
	// Contains explainability metrics for a model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_model_package#explainability SagemakerModelPackage#explainability}
	Explainability *SagemakerModelPackageDriftCheckBaselinesExplainability `field:"optional" json:"explainability" yaml:"explainability"`
	// Represents the drift check data quality baselines that can be used when the model monitor is set using the model package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_model_package#model_data_quality SagemakerModelPackage#model_data_quality}
	ModelDataQuality *SagemakerModelPackageDriftCheckBaselinesModelDataQuality `field:"optional" json:"modelDataQuality" yaml:"modelDataQuality"`
	// Represents the drift check model quality baselines that can be used when the model monitor is set using the model package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_model_package#model_quality SagemakerModelPackage#model_quality}
	ModelQuality *SagemakerModelPackageDriftCheckBaselinesModelQuality `field:"optional" json:"modelQuality" yaml:"modelQuality"`
}

