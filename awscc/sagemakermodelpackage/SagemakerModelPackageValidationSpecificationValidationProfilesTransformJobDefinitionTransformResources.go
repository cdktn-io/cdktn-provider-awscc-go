// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage


type SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformResources struct {
	// The number of ML compute instances to use in the transform job.
	//
	// For distributed transform jobs, specify a value greater than 1. The default value is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_model_package#instance_count SagemakerModelPackage#instance_count}
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// The ML compute instance type for the transform job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_model_package#instance_type SagemakerModelPackage#instance_type}
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt model data on the storage volume attached to the ML compute instance(s) that run the batch transform job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_model_package#volume_kms_key_id SagemakerModelPackage#volume_kms_key_id}
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
}

