// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage


type SagemakerModelPackageModelMetricsExplainability struct {
	// Represents a Metric Source Object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_model_package#report SagemakerModelPackage#report}
	Report *SagemakerModelPackageModelMetricsExplainabilityReport `field:"optional" json:"report" yaml:"report"`
}

