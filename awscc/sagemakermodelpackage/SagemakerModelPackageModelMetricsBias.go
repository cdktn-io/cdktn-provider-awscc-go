// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage


type SagemakerModelPackageModelMetricsBias struct {
	// Represents a Metric Source Object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_model_package#post_training_report SagemakerModelPackage#post_training_report}
	PostTrainingReport *SagemakerModelPackageModelMetricsBiasPostTrainingReport `field:"optional" json:"postTrainingReport" yaml:"postTrainingReport"`
	// Represents a Metric Source Object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_model_package#pre_training_report SagemakerModelPackage#pre_training_report}
	PreTrainingReport *SagemakerModelPackageModelMetricsBiasPreTrainingReport `field:"optional" json:"preTrainingReport" yaml:"preTrainingReport"`
	// Represents a Metric Source Object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_model_package#report SagemakerModelPackage#report}
	Report *SagemakerModelPackageModelMetricsBiasReport `field:"optional" json:"report" yaml:"report"`
}

