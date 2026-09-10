// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint


type DmsEndpointRedshiftSettings struct {
	// A value that indicates to allow any date format, including invalid formats such as 00/00/00 00:00:00, to be loaded without generating an error.
	//
	// You can choose true or false (the default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#accept_any_date DmsEndpoint#accept_any_date}
	AcceptAnyDate interface{} `field:"optional" json:"acceptAnyDate" yaml:"acceptAnyDate"`
	// Code to run after connecting.
	//
	// This parameter should contain the code itself, not the name of a file containing the code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#after_connect_script DmsEndpoint#after_connect_script}
	AfterConnectScript *string `field:"optional" json:"afterConnectScript" yaml:"afterConnectScript"`
	// An S3 folder where the comma-separated-value (.csv) files are stored before being uploaded to the target Redshift cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#bucket_folder DmsEndpoint#bucket_folder}
	BucketFolder *string `field:"optional" json:"bucketFolder" yaml:"bucketFolder"`
	// The name of the intermediate S3 bucket used to store .csv files before uploading data to Redshift.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#bucket_name DmsEndpoint#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// If Amazon Redshift is configured to support case sensitive schema names, set CaseSensitiveNames to true. The default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#case_sensitive_names DmsEndpoint#case_sensitive_names}
	CaseSensitiveNames interface{} `field:"optional" json:"caseSensitiveNames" yaml:"caseSensitiveNames"`
	// If you set CompUpdate to true Amazon Redshift applies automatic compression if the table is empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#comp_update DmsEndpoint#comp_update}
	CompUpdate interface{} `field:"optional" json:"compUpdate" yaml:"compUpdate"`
	// A value that sets the amount of time to wait (in milliseconds) before timing out, beginning from when you initially establish a connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#connection_timeout DmsEndpoint#connection_timeout}
	ConnectionTimeout *float64 `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// The date format that you are using.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#date_format DmsEndpoint#date_format}
	DateFormat *string `field:"optional" json:"dateFormat" yaml:"dateFormat"`
	// A value that specifies whether AWS DMS should migrate empty CHAR and VARCHAR fields as NULL.
	//
	// A value of true sets empty CHAR and VARCHAR fields to null. The default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#empty_as_null DmsEndpoint#empty_as_null}
	EmptyAsNull interface{} `field:"optional" json:"emptyAsNull" yaml:"emptyAsNull"`
	// The type of server-side encryption that you want to use for your data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#encryption_mode DmsEndpoint#encryption_mode}
	EncryptionMode *string `field:"optional" json:"encryptionMode" yaml:"encryptionMode"`
	// This setting is only valid for a full-load migration task.
	//
	// Set ExplicitIds to true to have tables with IDENTITY columns override their auto-generated values with explicit values loaded from the source data files used to populate the tables. The default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#explicit_ids DmsEndpoint#explicit_ids}
	ExplicitIds interface{} `field:"optional" json:"explicitIds" yaml:"explicitIds"`
	// The number of threads used to upload a single file.
	//
	// This parameter accepts a value from 1 through 64. It defaults to 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#file_transfer_upload_streams DmsEndpoint#file_transfer_upload_streams}
	FileTransferUploadStreams *float64 `field:"optional" json:"fileTransferUploadStreams" yaml:"fileTransferUploadStreams"`
	// The amount of time to wait (in milliseconds) before timing out of operations performed by AWS DMS on a Redshift cluster, such as Redshift COPY, INSERT, DELETE, and UPDATE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#load_timeout DmsEndpoint#load_timeout}
	LoadTimeout *float64 `field:"optional" json:"loadTimeout" yaml:"loadTimeout"`
	// When true, lets Redshift migrate the boolean type as boolean.
	//
	// By default, Redshift migrates booleans as varchar(1). You must set this setting on both the source and target endpoints for it to take effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#map_boolean_as_boolean DmsEndpoint#map_boolean_as_boolean}
	MapBooleanAsBoolean interface{} `field:"optional" json:"mapBooleanAsBoolean" yaml:"mapBooleanAsBoolean"`
	// The maximum size (in KB) of any .csv file used to load data on an S3 bucket and transfer data to Amazon Redshift. It defaults to 1048576KB (1 GB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#max_file_size DmsEndpoint#max_file_size}
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// A value that specifies to remove surrounding quotation marks from strings in the incoming data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#remove_quotes DmsEndpoint#remove_quotes}
	RemoveQuotes interface{} `field:"optional" json:"removeQuotes" yaml:"removeQuotes"`
	// A value that specifies to replaces the invalid characters specified in ReplaceInvalidChars, substituting the specified characters instead.
	//
	// The default is "?".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#replace_chars DmsEndpoint#replace_chars}
	ReplaceChars *string `field:"optional" json:"replaceChars" yaml:"replaceChars"`
	// A list of characters that you want to replace. Use with ReplaceChars.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#replace_invalid_chars DmsEndpoint#replace_invalid_chars}
	ReplaceInvalidChars *string `field:"optional" json:"replaceInvalidChars" yaml:"replaceInvalidChars"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in SecretsManagerSecret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#secrets_manager_access_role_arn DmsEndpoint#secrets_manager_access_role_arn}
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the SecretsManagerSecret that contains the Amazon Redshift endpoint connection details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#secrets_manager_secret_id DmsEndpoint#secrets_manager_secret_id}
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// The AWS KMS key ID. If you are using SSE_KMS for the EncryptionMode, provide this key ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#server_side_encryption_kms_key_id DmsEndpoint#server_side_encryption_kms_key_id}
	ServerSideEncryptionKmsKeyId *string `field:"optional" json:"serverSideEncryptionKmsKeyId" yaml:"serverSideEncryptionKmsKeyId"`
	// The Amazon Resource Name (ARN) of the IAM role that has access to the Amazon Redshift service.
	//
	// The role must allow the iam:PassRole action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#service_access_role_arn DmsEndpoint#service_access_role_arn}
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// The time format that you want to use. Valid values are auto (case-sensitive), 'timeformat_string', 'epochsecs', or 'epochmillisecs'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#time_format DmsEndpoint#time_format}
	TimeFormat *string `field:"optional" json:"timeFormat" yaml:"timeFormat"`
	// A value that specifies to remove the trailing white space characters from a VARCHAR string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#trim_blanks DmsEndpoint#trim_blanks}
	TrimBlanks interface{} `field:"optional" json:"trimBlanks" yaml:"trimBlanks"`
	// A value that specifies to truncate data in columns to the appropriate number of characters, so that the data fits in the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#truncate_columns DmsEndpoint#truncate_columns}
	TruncateColumns interface{} `field:"optional" json:"truncateColumns" yaml:"truncateColumns"`
	// The size (in KB) of the in-memory file write buffer used when generating .csv files on the local disk at the DMS replication instance. The default value is 1000 (buffer size is 1000KB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_endpoint#write_buffer_size DmsEndpoint#write_buffer_size}
	WriteBufferSize *float64 `field:"optional" json:"writeBufferSize" yaml:"writeBufferSize"`
}

