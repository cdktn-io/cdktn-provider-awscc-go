// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage


type SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinition struct {
	// A string that determines the number of records included in a single mini-batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_model_package#batch_strategy SagemakerModelPackage#batch_strategy}
	BatchStrategy *string `field:"optional" json:"batchStrategy" yaml:"batchStrategy"`
	// Sets the environment variables in the Docker container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_model_package#environment SagemakerModelPackage#environment}
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The maximum number of parallel requests that can be sent to each instance in a transform job.
	//
	// The default value is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_model_package#max_concurrent_transforms SagemakerModelPackage#max_concurrent_transforms}
	MaxConcurrentTransforms *float64 `field:"optional" json:"maxConcurrentTransforms" yaml:"maxConcurrentTransforms"`
	// The maximum payload size allowed, in MB. A payload is the data portion of a record (without metadata).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_model_package#max_payload_in_mb SagemakerModelPackage#max_payload_in_mb}
	MaxPayloadInMb *float64 `field:"optional" json:"maxPayloadInMb" yaml:"maxPayloadInMb"`
	// Describes the input source of a transform job and the way the transform job consumes it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_model_package#transform_input SagemakerModelPackage#transform_input}
	TransformInput *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInput `field:"optional" json:"transformInput" yaml:"transformInput"`
	// Describes the results of a transform job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_model_package#transform_output SagemakerModelPackage#transform_output}
	TransformOutput *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutput `field:"optional" json:"transformOutput" yaml:"transformOutput"`
	// Describes the resources, including ML instance types and ML instance count, to use for transform job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_model_package#transform_resources SagemakerModelPackage#transform_resources}
	TransformResources *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformResources `field:"optional" json:"transformResources" yaml:"transformResources"`
}

