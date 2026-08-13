// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage


type SagemakerModelPackageModelPackageStatusDetailsValidationStatuses struct {
	// If the overall status is Failed, the reason for the failure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_model_package#failure_reason SagemakerModelPackage#failure_reason}
	FailureReason *string `field:"optional" json:"failureReason" yaml:"failureReason"`
	// The name of the model package for which the overall status is being reported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_model_package#name SagemakerModelPackage#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The current status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_model_package#status SagemakerModelPackage#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

