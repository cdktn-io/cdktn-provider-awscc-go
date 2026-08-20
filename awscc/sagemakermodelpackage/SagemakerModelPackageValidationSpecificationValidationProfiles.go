// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage


type SagemakerModelPackageValidationSpecificationValidationProfiles struct {
	// The name of the profile for the model package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_model_package#profile_name SagemakerModelPackage#profile_name}
	ProfileName *string `field:"optional" json:"profileName" yaml:"profileName"`
	// Defines the input needed to run a transform job using the inference specification specified in the algorithm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_model_package#transform_job_definition SagemakerModelPackage#transform_job_definition}
	TransformJobDefinition *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinition `field:"optional" json:"transformJobDefinition" yaml:"transformJobDefinition"`
}

