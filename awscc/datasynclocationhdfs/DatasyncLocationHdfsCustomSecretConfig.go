// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationhdfs


type DatasyncLocationHdfsCustomSecretConfig struct {
	// Specifies the ARN for the AWS Identity and Access Management role that DataSync uses to access the secret specified for SecretArn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datasync_location_hdfs#secret_access_role_arn DatasyncLocationHdfs#secret_access_role_arn}
	SecretAccessRoleArn *string `field:"optional" json:"secretAccessRoleArn" yaml:"secretAccessRoleArn"`
	// Specifies the ARN for a customer created AWS Secrets Manager secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datasync_location_hdfs#secret_arn DatasyncLocationHdfs#secret_arn}
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

