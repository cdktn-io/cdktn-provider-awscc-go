// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage


type SagemakerModelPackageInferenceSpecificationContainersModelDataSource struct {
	// Specifies the S3 location of ML model data to deploy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_model_package#s3_data_source SagemakerModelPackage#s3_data_source}
	S3DataSource *SagemakerModelPackageInferenceSpecificationContainersModelDataSourceS3DataSource `field:"optional" json:"s3DataSource" yaml:"s3DataSource"`
}

