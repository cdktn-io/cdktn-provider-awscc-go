// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint


type DmsEndpointGcpMySqlSettings struct {
	// Specifies a script to run immediately after AWS DMS connects to the endpoint.
	//
	// The migration task continues running regardless if the SQL statement succeeds or fails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_endpoint#after_connect_script DmsEndpoint#after_connect_script}
	AfterConnectScript *string `field:"optional" json:"afterConnectScript" yaml:"afterConnectScript"`
	// Adjusts the behavior of AWS DMS when migrating from an SQL Server source database that is hosted as part of an Always On availability group cluster.
	//
	// If you need AWS DMS to poll all the nodes in the Always On cluster for transaction backups, set this attribute to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_endpoint#clean_source_metadata_on_mismatch DmsEndpoint#clean_source_metadata_on_mismatch}
	CleanSourceMetadataOnMismatch interface{} `field:"optional" json:"cleanSourceMetadataOnMismatch" yaml:"cleanSourceMetadataOnMismatch"`
	// Database name for the endpoint.
	//
	// For a MySQL source or target endpoint, don't explicitly specify the database using the DatabaseName request parameter on either the CreateEndpoint or ModifyEndpoint API call. Specifying DatabaseName when you create or modify a MySQL endpoint replicates all the task tables to this single database. For MySQL endpoints, you specify the database only when you specify the schema in the table-mapping rules of the AWS DMS task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_endpoint#database_name DmsEndpoint#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Specifies how often to check the binary log for new changes/events when the database is idle.
	//
	// The default is five seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_endpoint#events_poll_interval DmsEndpoint#events_poll_interval}
	EventsPollInterval *float64 `field:"optional" json:"eventsPollInterval" yaml:"eventsPollInterval"`
	// Specifies the maximum size (in KB) of any .csv file used to transfer data to a MySQL-compatible database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_endpoint#max_file_size DmsEndpoint#max_file_size}
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Improves performance when loading data into the MySQL-compatible target database.
	//
	// Specifies how many threads to use to load the data into the MySQL-compatible target database. Setting a large number of threads can have an adverse effect on database performance, because a separate connection is required for each thread. The default is one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_endpoint#parallel_load_threads DmsEndpoint#parallel_load_threads}
	ParallelLoadThreads *float64 `field:"optional" json:"parallelLoadThreads" yaml:"parallelLoadThreads"`
	// Endpoint connection password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_endpoint#password DmsEndpoint#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The port used by the endpoint database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_endpoint#port DmsEndpoint#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in SecretsManagerSecret.
	//
	// The role must allow the iam:PassRole action. SecretsManagerSecret has the value of the AWS Secrets Manager secret that allows access to the MySQL endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_endpoint#secrets_manager_access_role_arn DmsEndpoint#secrets_manager_access_role_arn}
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the SecretsManagerSecret that contains the MySQL endpoint connection details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_endpoint#secrets_manager_secret_id DmsEndpoint#secrets_manager_secret_id}
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// The MySQL host name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_endpoint#server_name DmsEndpoint#server_name}
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// Specifies the time zone for the source MySQL database. Don't enclose time zones in single quotation marks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_endpoint#server_timezone DmsEndpoint#server_timezone}
	ServerTimezone *string `field:"optional" json:"serverTimezone" yaml:"serverTimezone"`
	// Specifies the time zone for the source MySQL database. Don't enclose time zones in single quotation marks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dms_endpoint#username DmsEndpoint#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

