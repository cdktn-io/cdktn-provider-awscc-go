// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint


type DmsEndpointIbmDb2Settings struct {
	// For ongoing replication (CDC), use CurrentLSN to specify a log sequence number (LSN) where you want the replication to start.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#current_lsn DmsEndpoint#current_lsn}
	CurrentLsn *string `field:"optional" json:"currentLsn" yaml:"currentLsn"`
	// If true, AWS DMS saves any .csv files to the Db2 LUW target that were used to replicate data. DMS uses these files for analysis and troubleshooting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#keep_csv_files DmsEndpoint#keep_csv_files}
	KeepCsvFiles interface{} `field:"optional" json:"keepCsvFiles" yaml:"keepCsvFiles"`
	// The amount of time (in milliseconds) before AWS DMS times out operations performed by DMS on the Db2 target.
	//
	// The default value is 1200 (20 minutes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#load_timeout DmsEndpoint#load_timeout}
	LoadTimeout *float64 `field:"optional" json:"loadTimeout" yaml:"loadTimeout"`
	// Specifies the maximum size (in KB) of .csv files used to transfer data to Db2 LUW.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#max_file_size DmsEndpoint#max_file_size}
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Maximum number of bytes per read, as a NUMBER value. The default is 64 KB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#max_k_bytes_per_read DmsEndpoint#max_k_bytes_per_read}
	MaxKBytesPerRead *float64 `field:"optional" json:"maxKBytesPerRead" yaml:"maxKBytesPerRead"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in SecretsManagerSecret.
	//
	// The role must allow the iam:PassRole action. SecretsManagerSecret has the value ofthe AWS Secrets Manager secret that allows access to the Db2 LUW endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#secrets_manager_access_role_arn DmsEndpoint#secrets_manager_access_role_arn}
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the SecretsManagerSecret that contains the IBMDB2 endpoint connection details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#secrets_manager_secret_id DmsEndpoint#secrets_manager_secret_id}
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// Enables ongoing replication (CDC) as a BOOLEAN value. The default is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#set_data_capture_changes DmsEndpoint#set_data_capture_changes}
	SetDataCaptureChanges interface{} `field:"optional" json:"setDataCaptureChanges" yaml:"setDataCaptureChanges"`
	// The size (in KB) of the in-memory file write buffer used when generating .csv files on the local disk on the DMS replication instance. The default value is 1024 (1 MB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#write_buffer_size DmsEndpoint#write_buffer_size}
	WriteBufferSize *float64 `field:"optional" json:"writeBufferSize" yaml:"writeBufferSize"`
}

