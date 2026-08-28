// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodel


type SagemakerModelContainers struct {
	// This parameter is ignored for models that contain only a PrimaryContainer.
	//
	// When a ContainerDefinition is part of an inference pipeline, the value of the parameter uniquely identifies the container for the purposes of logging and metrics. For information, see [Use Logs and Metrics to Monitor an Inference Pipeline](https://docs.aws.amazon.com/sagemaker/latest/dg/inference-pipeline-logs-metrics.html). If you don't specify a value for this parameter for a ContainerDefinition that is part of an inference pipeline, a unique name is automatically assigned based on the position of the ContainerDefinition in the pipeline. If you specify a value for the ContainerHostName for any ContainerDefinition that is part of an inference pipeline, you must specify a value for the ContainerHostName parameter of every ContainerDefinition in that pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model#container_hostname SagemakerModel#container_hostname}
	ContainerHostname *string `field:"optional" json:"containerHostname" yaml:"containerHostname"`
	// The environment variables to set in the Docker container. Don't include any sensitive data in your environment variables.
	//
	// The maximum length of each key and value in the Environment map is 1024 bytes. The maximum length of all keys and values in the map, combined, is 32 KB. If you pass multiple containers to a CreateModel request, then the maximum length of all of their maps, combined, is also 32 KB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model#environment SagemakerModel#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// The path where inference code is stored.
	//
	// This can be either in Amazon EC2 Container Registry or in a Docker registry that is accessible from the same VPC that you configure for your endpoint. If you are using your own custom algorithm instead of an algorithm provided by SageMaker, the inference code must meet SageMaker requirements. SageMaker supports both registry/repository[:tag] and registry/repository[@digest] image path formats. For more information, see [Using Your Own Algorithms with Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model#image SagemakerModel#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model#image_config SagemakerModel#image_config}
	ImageConfig *SagemakerModelContainersImageConfig `field:"optional" json:"imageConfig" yaml:"imageConfig"`
	// The inference specification name in the model package version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model#inference_specification_name SagemakerModel#inference_specification_name}
	InferenceSpecificationName *string `field:"optional" json:"inferenceSpecificationName" yaml:"inferenceSpecificationName"`
	// Whether the container hosts a single model or multiple models.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model#mode SagemakerModel#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Specifies the location of ML model data to deploy.
	//
	// If specified, you must specify one and only one of the available data sources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model#model_data_source SagemakerModel#model_data_source}
	ModelDataSource *SagemakerModelContainersModelDataSource `field:"optional" json:"modelDataSource" yaml:"modelDataSource"`
	// The S3 path where the model artifacts, which result from model training, are stored.
	//
	// This path must point to a single gzip compressed tar archive (.tar.gz suffix). The S3 path is required for SageMaker built-in algorithms, but not if you use your own algorithms. For more information on built-in algorithms, see [Common Parameters](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-algo-docker-registry-paths.html).
	//
	// If you provide a value for this parameter, SageMaker uses AWS Security Token Service to download model artifacts from the S3 path you provide. AWS STS is activated in your AWS account by default. If you previously deactivated AWS STS for a region, you need to reactivate AWS STS for that region. For more information, see [Activating and Deactivating AWS STS in an AWS Region](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html) in the AWS Identity and Access Management User Guide
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model#model_data_url SagemakerModel#model_data_url}
	ModelDataUrl *string `field:"optional" json:"modelDataUrl" yaml:"modelDataUrl"`
	// The name or Amazon Resource Name (ARN) of the model package to use to create the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model#model_package_name SagemakerModel#model_package_name}
	ModelPackageName *string `field:"optional" json:"modelPackageName" yaml:"modelPackageName"`
	// Specifies additional configuration for multi-model endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model#multi_model_config SagemakerModel#multi_model_config}
	MultiModelConfig *SagemakerModelContainersMultiModelConfig `field:"optional" json:"multiModelConfig" yaml:"multiModelConfig"`
}

