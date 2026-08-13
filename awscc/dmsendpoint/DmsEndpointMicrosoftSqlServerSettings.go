// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint


type DmsEndpointMicrosoftSqlServerSettings struct {
	// The maximum size of the packets (in bytes) used to transfer data using BCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dms_endpoint#bcp_packet_size DmsEndpoint#bcp_packet_size}
	BcpPacketSize *float64 `field:"optional" json:"bcpPacketSize" yaml:"bcpPacketSize"`
	// Specifies a file group for the AWS DMS internal tables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dms_endpoint#control_tables_file_group DmsEndpoint#control_tables_file_group}
	ControlTablesFileGroup *string `field:"optional" json:"controlTablesFileGroup" yaml:"controlTablesFileGroup"`
	// Database name for the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dms_endpoint#database_name DmsEndpoint#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Forces LOB lookup on inline LOB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dms_endpoint#force_lob_lookup DmsEndpoint#force_lob_lookup}
	ForceLobLookup interface{} `field:"optional" json:"forceLobLookup" yaml:"forceLobLookup"`
	// Endpoint connection password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dms_endpoint#password DmsEndpoint#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Endpoint TCP port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dms_endpoint#port DmsEndpoint#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Cleans and recreates table metadata information on the replication instance when a mismatch occurs.
	//
	// An example is a situation where running an alter DDL statement on a table might result in different information about the table cached in the replication instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dms_endpoint#query_single_always_on_node DmsEndpoint#query_single_always_on_node}
	QuerySingleAlwaysOnNode interface{} `field:"optional" json:"querySingleAlwaysOnNode" yaml:"querySingleAlwaysOnNode"`
	// When this attribute is set to Y, AWS DMS only reads changes from transaction log backups and doesn't read from the active transaction log file during ongoing replication.
	//
	// Setting this parameter to Y enables you to control active transaction log file growth during full load and ongoing replication tasks. However, it can add some source latency to ongoing replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dms_endpoint#read_backup_only DmsEndpoint#read_backup_only}
	ReadBackupOnly interface{} `field:"optional" json:"readBackupOnly" yaml:"readBackupOnly"`
	// Use this attribute to minimize the need to access the backup log and enable AWS DMS to prevent truncation using one of the following two methods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dms_endpoint#safeguard_policy DmsEndpoint#safeguard_policy}
	SafeguardPolicy *string `field:"optional" json:"safeguardPolicy" yaml:"safeguardPolicy"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in SecretsManagerSecret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dms_endpoint#secrets_manager_access_role_arn DmsEndpoint#secrets_manager_access_role_arn}
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the SecretsManagerSecret that contains the MicrosoftSQLServer endpoint connection details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dms_endpoint#secrets_manager_secret_id DmsEndpoint#secrets_manager_secret_id}
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// Fully qualified domain name of the endpoint.
	//
	// For an Amazon RDS SQL Server instance, this is the output of DescribeDBInstances, in the Endpoint.Address field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dms_endpoint#server_name DmsEndpoint#server_name}
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// Indicates the mode used to fetch CDC data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dms_endpoint#tlog_access_mode DmsEndpoint#tlog_access_mode}
	TlogAccessMode *string `field:"optional" json:"tlogAccessMode" yaml:"tlogAccessMode"`
	// Use the TrimSpaceInChar source endpoint setting to right-trim data on CHAR and NCHAR data types during migration.
	//
	// Setting TrimSpaceInChar does not left-trim data. The default value is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dms_endpoint#trim_space_in_char DmsEndpoint#trim_space_in_char}
	TrimSpaceInChar interface{} `field:"optional" json:"trimSpaceInChar" yaml:"trimSpaceInChar"`
	// Use this to attribute to transfer data for full-load operations using BCP.
	//
	// When the target table contains an identity column that does not exist in the source table, you must disable the use BCP for loading table option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dms_endpoint#use_bcp_full_load DmsEndpoint#use_bcp_full_load}
	UseBcpFullLoad interface{} `field:"optional" json:"useBcpFullLoad" yaml:"useBcpFullLoad"`
	// Endpoint connection user name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dms_endpoint#username DmsEndpoint#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
	// When this attribute is set to Y, DMS processes third-party transaction log backups if they are created in native format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dms_endpoint#use_third_party_backup_device DmsEndpoint#use_third_party_backup_device}
	UseThirdPartyBackupDevice interface{} `field:"optional" json:"useThirdPartyBackupDevice" yaml:"useThirdPartyBackupDevice"`
}

