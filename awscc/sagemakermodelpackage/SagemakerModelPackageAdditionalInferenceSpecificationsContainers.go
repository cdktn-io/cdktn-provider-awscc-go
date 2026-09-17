// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage


type SagemakerModelPackageAdditionalInferenceSpecificationsContainers struct {
	// The DNS host name for the Docker container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_package#container_hostname SagemakerModelPackage#container_hostname}
	ContainerHostname *string `field:"optional" json:"containerHostname" yaml:"containerHostname"`
	// Sets the environment variables in the Docker container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_package#environment SagemakerModelPackage#environment}
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The machine learning framework of the model package container image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_package#framework SagemakerModelPackage#framework}
	Framework *string `field:"optional" json:"framework" yaml:"framework"`
	// The framework version of the Model Package Container Image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_package#framework_version SagemakerModelPackage#framework_version}
	FrameworkVersion *string `field:"optional" json:"frameworkVersion" yaml:"frameworkVersion"`
	// The Amazon EC2 Container Registry (Amazon ECR) path where inference code is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_package#image SagemakerModelPackage#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// An MD5 hash of the training algorithm that identifies the Docker image used for training.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_package#image_digest SagemakerModelPackage#image_digest}
	ImageDigest *string `field:"optional" json:"imageDigest" yaml:"imageDigest"`
	// Specifies the location of ML model data to deploy during endpoint creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_package#model_data_source SagemakerModelPackage#model_data_source}
	ModelDataSource *SagemakerModelPackageAdditionalInferenceSpecificationsContainersModelDataSource `field:"optional" json:"modelDataSource" yaml:"modelDataSource"`
	// A structure with Model Input details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_package#model_data_url SagemakerModelPackage#model_data_url}
	ModelDataUrl *string `field:"optional" json:"modelDataUrl" yaml:"modelDataUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_package#model_input SagemakerModelPackage#model_input}.
	ModelInput *SagemakerModelPackageAdditionalInferenceSpecificationsContainersModelInput `field:"optional" json:"modelInput" yaml:"modelInput"`
	// The name of a pre-trained machine learning benchmarked by Amazon SageMaker Inference Recommender model that matches your model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_package#nearest_model_name SagemakerModelPackage#nearest_model_name}
	NearestModelName *string `field:"optional" json:"nearestModelName" yaml:"nearestModelName"`
}

