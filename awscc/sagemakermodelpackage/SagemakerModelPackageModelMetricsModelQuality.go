// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage


type SagemakerModelPackageModelMetricsModelQuality struct {
	// Represents a Metric Source Object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_package#constraints SagemakerModelPackage#constraints}
	Constraints *SagemakerModelPackageModelMetricsModelQualityConstraints `field:"optional" json:"constraints" yaml:"constraints"`
	// Represents a Metric Source Object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_package#statistics SagemakerModelPackage#statistics}
	Statistics *SagemakerModelPackageModelMetricsModelQualityStatistics `field:"optional" json:"statistics" yaml:"statistics"`
}

