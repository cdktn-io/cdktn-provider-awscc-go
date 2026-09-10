// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage


type SagemakerModelPackageAdditionalInferenceSpecificationsToAdd struct {
	// The Amazon ECR registry path of the Docker image that contains the inference code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_model_package#containers SagemakerModelPackage#containers}
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// A description of the additional Inference specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_model_package#description SagemakerModelPackage#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A unique name to identify the additional inference specification.
	//
	// The name must be unique within the list of your additional inference specifications for a particular model package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_model_package#name SagemakerModelPackage#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The supported MIME types for the input data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_model_package#supported_content_types SagemakerModelPackage#supported_content_types}
	SupportedContentTypes *[]*string `field:"optional" json:"supportedContentTypes" yaml:"supportedContentTypes"`
	// A list of the instance types that are used to generate inferences in real-time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_model_package#supported_realtime_inference_instance_types SagemakerModelPackage#supported_realtime_inference_instance_types}
	SupportedRealtimeInferenceInstanceTypes *[]*string `field:"optional" json:"supportedRealtimeInferenceInstanceTypes" yaml:"supportedRealtimeInferenceInstanceTypes"`
	// The supported MIME types for the output data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_model_package#supported_response_mime_types SagemakerModelPackage#supported_response_mime_types}
	SupportedResponseMimeTypes *[]*string `field:"optional" json:"supportedResponseMimeTypes" yaml:"supportedResponseMimeTypes"`
	// A list of the instance types on which a transformation job can be run or on which an endpoint can be deployed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_model_package#supported_transform_instance_types SagemakerModelPackage#supported_transform_instance_types}
	SupportedTransformInstanceTypes *[]*string `field:"optional" json:"supportedTransformInstanceTypes" yaml:"supportedTransformInstanceTypes"`
}

