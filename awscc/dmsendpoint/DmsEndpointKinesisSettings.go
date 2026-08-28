// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint


type DmsEndpointKinesisSettings struct {
	// Shows detailed control information for table definition, column definition, and table and column changes in the Kinesis message output.
	//
	// The default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_endpoint#include_control_details DmsEndpoint#include_control_details}
	IncludeControlDetails interface{} `field:"optional" json:"includeControlDetails" yaml:"includeControlDetails"`
	// Include NULL and empty columns for records migrated to the endpoint. The default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_endpoint#include_null_and_empty DmsEndpoint#include_null_and_empty}
	IncludeNullAndEmpty interface{} `field:"optional" json:"includeNullAndEmpty" yaml:"includeNullAndEmpty"`
	// Shows the partition value within the Kinesis message output, unless the partition type is schema-table-type. The default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_endpoint#include_partition_value DmsEndpoint#include_partition_value}
	IncludePartitionValue interface{} `field:"optional" json:"includePartitionValue" yaml:"includePartitionValue"`
	// Includes any data definition language (DDL) operations that change the table in the control data, such as rename-table, drop-table, add-column, drop-column, and rename-column.
	//
	// The default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_endpoint#include_table_alter_operations DmsEndpoint#include_table_alter_operations}
	IncludeTableAlterOperations interface{} `field:"optional" json:"includeTableAlterOperations" yaml:"includeTableAlterOperations"`
	// Provides detailed transaction information from the source database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_endpoint#include_transaction_details DmsEndpoint#include_transaction_details}
	IncludeTransactionDetails interface{} `field:"optional" json:"includeTransactionDetails" yaml:"includeTransactionDetails"`
	// The output format for the records created on the endpoint.
	//
	// The message format is JSON (default) or JSON_UNFORMATTED (a single line with no tab).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_endpoint#message_format DmsEndpoint#message_format}
	MessageFormat *string `field:"optional" json:"messageFormat" yaml:"messageFormat"`
	// Set this optional parameter to true to avoid adding a '0x' prefix to raw data in hexadecimal format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_endpoint#no_hex_prefix DmsEndpoint#no_hex_prefix}
	NoHexPrefix interface{} `field:"optional" json:"noHexPrefix" yaml:"noHexPrefix"`
	// Prefixes schema and table names to partition values, when the partition type is primary-key-type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_endpoint#partition_include_schema_table DmsEndpoint#partition_include_schema_table}
	PartitionIncludeSchemaTable interface{} `field:"optional" json:"partitionIncludeSchemaTable" yaml:"partitionIncludeSchemaTable"`
	// The Amazon Resource Name (ARN) for the IAM role that AWS DMS uses to write to the Kinesis data stream.
	//
	// The role must allow the iam:PassRole action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_endpoint#service_access_role_arn DmsEndpoint#service_access_role_arn}
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// The Amazon Resource Name (ARN) for the Amazon Kinesis Data Streams endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_endpoint#stream_arn DmsEndpoint#stream_arn}
	StreamArn *string `field:"optional" json:"streamArn" yaml:"streamArn"`
}

