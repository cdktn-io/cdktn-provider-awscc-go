// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage


type SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputDataSource struct {
	// Describes the S3 data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_model_package#s3_data_source SagemakerModelPackage#s3_data_source}
	S3DataSource *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputDataSourceS3DataSource `field:"optional" json:"s3DataSource" yaml:"s3DataSource"`
}

