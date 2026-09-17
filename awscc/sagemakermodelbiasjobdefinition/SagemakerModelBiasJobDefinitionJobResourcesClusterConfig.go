// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelbiasjobdefinition


type SagemakerModelBiasJobDefinitionJobResourcesClusterConfig struct {
	// The number of ML compute instances to use in the model monitoring job.
	//
	// For distributed processing jobs, specify a value greater than 1. The default value is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_bias_job_definition#instance_count SagemakerModelBiasJobDefinition#instance_count}
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// The ML compute instance type for the processing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_bias_job_definition#instance_type SagemakerModelBiasJobDefinition#instance_type}
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The size of the ML storage volume, in gigabytes, that you want to provision.
	//
	// You must specify sufficient ML storage for your scenario.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_bias_job_definition#volume_size_in_gb SagemakerModelBiasJobDefinition#volume_size_in_gb}
	VolumeSizeInGb *float64 `field:"required" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance(s) that run the model monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_bias_job_definition#volume_kms_key_id SagemakerModelBiasJobDefinition#volume_kms_key_id}
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
}

