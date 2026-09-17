// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint


type DmsEndpointPostgreSqlSettings struct {
	// For use with change data capture (CDC) only, this attribute has AWS DMS bypass foreign keys and user triggers to reduce the time it takes to bulk load data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#after_connect_script DmsEndpoint#after_connect_script}
	AfterConnectScript *string `field:"optional" json:"afterConnectScript" yaml:"afterConnectScript"`
	// The Babelfish for Aurora PostgreSQL database name for the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#babelfish_database_name DmsEndpoint#babelfish_database_name}
	BabelfishDatabaseName *string `field:"optional" json:"babelfishDatabaseName" yaml:"babelfishDatabaseName"`
	// To capture DDL events, AWS DMS creates various artifacts in the PostgreSQL database when the task starts.
	//
	// You can later remove these artifacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#capture_ddls DmsEndpoint#capture_ddls}
	CaptureDdls interface{} `field:"optional" json:"captureDdls" yaml:"captureDdls"`
	// Specifies the default behavior of the replication's handling of PostgreSQL- compatible endpoints that require some additional configuration, such as Babelfish endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#database_mode DmsEndpoint#database_mode}
	DatabaseMode *string `field:"optional" json:"databaseMode" yaml:"databaseMode"`
	// The schema in which the operational DDL database artifacts are created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#ddl_artifacts_schema DmsEndpoint#ddl_artifacts_schema}
	DdlArtifactsSchema *string `field:"optional" json:"ddlArtifactsSchema" yaml:"ddlArtifactsSchema"`
	// Sets the client statement timeout for the PostgreSQL instance, in seconds. The default value is 60 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#execute_timeout DmsEndpoint#execute_timeout}
	ExecuteTimeout *float64 `field:"optional" json:"executeTimeout" yaml:"executeTimeout"`
	// When set to true, this value causes a task to fail if the actual size of a LOB column is greater than the specified LobMaxSize.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#fail_tasks_on_lob_truncation DmsEndpoint#fail_tasks_on_lob_truncation}
	FailTasksOnLobTruncation interface{} `field:"optional" json:"failTasksOnLobTruncation" yaml:"failTasksOnLobTruncation"`
	// The write-ahead log (WAL) heartbeat feature mimics a dummy transaction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#heartbeat_enable DmsEndpoint#heartbeat_enable}
	HeartbeatEnable interface{} `field:"optional" json:"heartbeatEnable" yaml:"heartbeatEnable"`
	// Sets the WAL heartbeat frequency (in minutes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#heartbeat_frequency DmsEndpoint#heartbeat_frequency}
	HeartbeatFrequency *float64 `field:"optional" json:"heartbeatFrequency" yaml:"heartbeatFrequency"`
	// Sets the schema in which the heartbeat artifacts are created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#heartbeat_schema DmsEndpoint#heartbeat_schema}
	HeartbeatSchema *string `field:"optional" json:"heartbeatSchema" yaml:"heartbeatSchema"`
	// When true, lets PostgreSQL migrate the boolean type as boolean.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#map_boolean_as_boolean DmsEndpoint#map_boolean_as_boolean}
	MapBooleanAsBoolean interface{} `field:"optional" json:"mapBooleanAsBoolean" yaml:"mapBooleanAsBoolean"`
	// Specifies the maximum size (in KB) of any .csv file used to transfer data to PostgreSQL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#max_file_size DmsEndpoint#max_file_size}
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Specifies the plugin to use to create a replication slot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#plugin_name DmsEndpoint#plugin_name}
	PluginName *string `field:"optional" json:"pluginName" yaml:"pluginName"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in SecretsManagerSecret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#secrets_manager_access_role_arn DmsEndpoint#secrets_manager_access_role_arn}
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the SecretsManagerSecret that contains the PostgreSQL endpoint connection details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#secrets_manager_secret_id DmsEndpoint#secrets_manager_secret_id}
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// Sets the name of a previously created logical replication slot for a change data capture (CDC) load of the PostgreSQL source instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#slot_name DmsEndpoint#slot_name}
	SlotName *string `field:"optional" json:"slotName" yaml:"slotName"`
}

