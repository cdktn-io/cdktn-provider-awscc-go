// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonedatasource


type DatazoneDataSourceConfigurationRedshiftRunConfigurationRedshiftCredentialConfiguration struct {
	// The ARN of a secret manager for an Amazon Redshift cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datazone_data_source#secret_manager_arn DatazoneDataSource#secret_manager_arn}
	SecretManagerArn *string `field:"optional" json:"secretManagerArn" yaml:"secretManagerArn"`
}

