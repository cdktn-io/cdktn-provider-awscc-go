// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodel


type SagemakerModelPrimaryContainerModelDataSourceS3DataSource struct {
	// Specifies how the ML model data is prepared.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_model#compression_type SagemakerModel#compression_type}
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Configuration information specifying which hub contents have accessible deployment options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_model#hub_access_config SagemakerModel#hub_access_config}
	HubAccessConfig *SagemakerModelPrimaryContainerModelDataSourceS3DataSourceHubAccessConfig `field:"optional" json:"hubAccessConfig" yaml:"hubAccessConfig"`
	// The access configuration file to control access to the ML model.
	//
	// You can explicitly accept the model end-user license agreement (EULA) within the `ModelAccessConfig`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_model#model_access_config SagemakerModel#model_access_config}
	ModelAccessConfig *SagemakerModelPrimaryContainerModelDataSourceS3DataSourceModelAccessConfig `field:"optional" json:"modelAccessConfig" yaml:"modelAccessConfig"`
	// Specifies the type of ML model data to deploy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_model#s3_data_type SagemakerModel#s3_data_type}
	S3DataType *string `field:"optional" json:"s3DataType" yaml:"s3DataType"`
	// Specifies the S3 path of ML model data to deploy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_model#s3_uri SagemakerModel#s3_uri}
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

