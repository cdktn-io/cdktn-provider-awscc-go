// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package forecastdataset


type ForecastDatasetEncryptionConfig struct {
	// KMS key used to encrypt the Dataset data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/forecast_dataset#kms_key_arn ForecastDataset#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The ARN of the IAM role that Amazon Forecast can assume to access the AWS KMS key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/forecast_dataset#role_arn ForecastDataset#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

