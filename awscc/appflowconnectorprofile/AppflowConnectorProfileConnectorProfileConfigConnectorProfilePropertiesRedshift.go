// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesRedshift struct {
	// The name of the Amazon S3 bucket associated with Redshift.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appflow_connector_profile#bucket_name AppflowConnectorProfile#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The object key for the destination bucket in which Amazon AppFlow will place the ?les.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appflow_connector_profile#bucket_prefix AppflowConnectorProfile#bucket_prefix}
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// The unique identifier of the Amazon Redshift cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appflow_connector_profile#cluster_identifier AppflowConnectorProfile#cluster_identifier}
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// The Amazon Resource Name (ARN) of the IAM role that grants Amazon AppFlow access to the data through the Amazon Redshift Data API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appflow_connector_profile#data_api_role_arn AppflowConnectorProfile#data_api_role_arn}
	DataApiRoleArn *string `field:"optional" json:"dataApiRoleArn" yaml:"dataApiRoleArn"`
	// The name of the Amazon Redshift database that will store the transferred data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appflow_connector_profile#database_name AppflowConnectorProfile#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The JDBC URL of the Amazon Redshift cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appflow_connector_profile#database_url AppflowConnectorProfile#database_url}
	DatabaseUrl *string `field:"optional" json:"databaseUrl" yaml:"databaseUrl"`
	// If Amazon AppFlow will connect to Amazon Redshift Serverless or Amazon Redshift cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appflow_connector_profile#is_redshift_serverless AppflowConnectorProfile#is_redshift_serverless}
	IsRedshiftServerless interface{} `field:"optional" json:"isRedshiftServerless" yaml:"isRedshiftServerless"`
	// The Amazon Resource Name (ARN) of the IAM role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appflow_connector_profile#role_arn AppflowConnectorProfile#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The name of the Amazon Redshift serverless workgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appflow_connector_profile#workgroup_name AppflowConnectorProfile#workgroup_name}
	WorkgroupName *string `field:"optional" json:"workgroupName" yaml:"workgroupName"`
}

