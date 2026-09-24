// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage


type SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersModelDataSourceS3DataSource struct {
	// Specifies how the ML model data is prepared.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_model_package#compression_type SagemakerModelPackage#compression_type}
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Specifies the access configuration file for the ML model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_model_package#model_access_config SagemakerModelPackage#model_access_config}
	ModelAccessConfig *SagemakerModelPackageAdditionalInferenceSpecificationsToAddContainersModelDataSourceS3DataSourceModelAccessConfig `field:"optional" json:"modelAccessConfig" yaml:"modelAccessConfig"`
	// Specifies the type of ML model data to deploy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_model_package#s3_data_type SagemakerModelPackage#s3_data_type}
	S3DataType *string `field:"optional" json:"s3DataType" yaml:"s3DataType"`
	// Specifies the S3 path of ML model data to deploy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_model_package#s3_uri SagemakerModelPackage#s3_uri}
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

