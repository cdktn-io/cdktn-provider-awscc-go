// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreamtable

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TimestreamTableConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name for the database which the table to be created belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/timestream_table#database_name TimestreamTable#database_name}
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The properties that determine whether magnetic store writes are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/timestream_table#magnetic_store_write_properties TimestreamTable#magnetic_store_write_properties}
	MagneticStoreWriteProperties *TimestreamTableMagneticStoreWriteProperties `field:"optional" json:"magneticStoreWriteProperties" yaml:"magneticStoreWriteProperties"`
	// The retention duration of the memory store and the magnetic store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/timestream_table#retention_properties TimestreamTable#retention_properties}
	RetentionProperties *TimestreamTableRetentionProperties `field:"optional" json:"retentionProperties" yaml:"retentionProperties"`
	// A Schema specifies the expected data model of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/timestream_table#schema TimestreamTable#schema}
	Schema *TimestreamTableSchema `field:"optional" json:"schema" yaml:"schema"`
	// The name for the table.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/timestream_table#table_name TimestreamTable#table_name}
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/timestream_table#tags TimestreamTable#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

