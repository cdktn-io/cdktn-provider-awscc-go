// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsdbproxy


type RdsDbProxyAuth struct {
	// The type of authentication that the proxy uses for connections from the proxy to the underlying database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/rds_db_proxy#auth_scheme RdsDbProxy#auth_scheme}
	AuthScheme *string `field:"optional" json:"authScheme" yaml:"authScheme"`
	// The type of authentication the proxy uses for connections from clients.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/rds_db_proxy#client_password_auth_type RdsDbProxy#client_password_auth_type}
	ClientPasswordAuthType *string `field:"optional" json:"clientPasswordAuthType" yaml:"clientPasswordAuthType"`
	// A user-specified description about the authentication used by a proxy to log in as a specific database user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/rds_db_proxy#description RdsDbProxy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether to require or disallow Amazon Web Services Identity and Access Management (IAM) authentication for connections to the proxy.
	//
	// The ENABLED value is valid only for proxies with RDS for Microsoft SQL Server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/rds_db_proxy#iam_auth RdsDbProxy#iam_auth}
	IamAuth *string `field:"optional" json:"iamAuth" yaml:"iamAuth"`
	// The Amazon Resource Name (ARN) representing the secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster.
	//
	// These secrets are stored within Amazon Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/rds_db_proxy#secret_arn RdsDbProxy#secret_arn}
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

