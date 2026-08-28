// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint


type DmsEndpointMySqlSettings struct {
	// Specifies a script to run immediately after AWS DMS connects to the endpoint.
	//
	// The migration task continues running regardless if the SQL statement succeeds or fails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_endpoint#after_connect_script DmsEndpoint#after_connect_script}
	AfterConnectScript *string `field:"optional" json:"afterConnectScript" yaml:"afterConnectScript"`
	// Cleans and recreates table metadata information on the replication instance when a mismatch occurs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_endpoint#clean_source_metadata_on_mismatch DmsEndpoint#clean_source_metadata_on_mismatch}
	CleanSourceMetadataOnMismatch interface{} `field:"optional" json:"cleanSourceMetadataOnMismatch" yaml:"cleanSourceMetadataOnMismatch"`
	// Specifies how often to check the binary log for new changes/events when the database is idle.
	//
	// The default is five seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_endpoint#events_poll_interval DmsEndpoint#events_poll_interval}
	EventsPollInterval *float64 `field:"optional" json:"eventsPollInterval" yaml:"eventsPollInterval"`
	// Specifies the maximum size (in KB) of any .csv file used to transfer data to a MySQL-compatible database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_endpoint#max_file_size DmsEndpoint#max_file_size}
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Improves performance when loading data into the MySQL-compatible target database.
	//
	// Specifies how many threads to use to load the data into the MySQL-compatible target database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_endpoint#parallel_load_threads DmsEndpoint#parallel_load_threads}
	ParallelLoadThreads *float64 `field:"optional" json:"parallelLoadThreads" yaml:"parallelLoadThreads"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in SecretsManagerSecret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_endpoint#secrets_manager_access_role_arn DmsEndpoint#secrets_manager_access_role_arn}
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the SecretsManagerSecret that contains the MySQL endpoint connection details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_endpoint#secrets_manager_secret_id DmsEndpoint#secrets_manager_secret_id}
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// Specifies the time zone for the source MySQL database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_endpoint#server_timezone DmsEndpoint#server_timezone}
	ServerTimezone *string `field:"optional" json:"serverTimezone" yaml:"serverTimezone"`
	// Specifies where to migrate source tables on the target, either to a single database or multiple databases.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_endpoint#target_db_type DmsEndpoint#target_db_type}
	TargetDbType *string `field:"optional" json:"targetDbType" yaml:"targetDbType"`
}

