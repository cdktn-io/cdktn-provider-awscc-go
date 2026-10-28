// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable


type GlueTableTableInputStorageDescriptor struct {
	// A list of reducer grouping columns, clustering columns, and bucketing columns in the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#bucket_columns GlueTable#bucket_columns}
	BucketColumns *[]*string `field:"optional" json:"bucketColumns" yaml:"bucketColumns"`
	// A list of the Columns in the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#columns GlueTable#columns}
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// True if the data in the table is compressed, or False if not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#compressed GlueTable#compressed}
	Compressed interface{} `field:"optional" json:"compressed" yaml:"compressed"`
	// The input format: SequenceFileInputFormat (binary), or TextInputFormat, or a custom format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#input_format GlueTable#input_format}
	InputFormat *string `field:"optional" json:"inputFormat" yaml:"inputFormat"`
	// The physical location of the table.
	//
	// By default, this takes the form of the warehouse location, followed by the database location in the warehouse, followed by the table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#location GlueTable#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Must be specified if the table contains any dimension columns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#number_of_buckets GlueTable#number_of_buckets}
	NumberOfBuckets *float64 `field:"optional" json:"numberOfBuckets" yaml:"numberOfBuckets"`
	// The output format: SequenceFileOutputFormat (binary), or IgnoreKeyTextOutputFormat, or a custom format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#output_format GlueTable#output_format}
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// The user-supplied properties in key-value form.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#parameters GlueTable#parameters}
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
	// An object that references a schema stored in the AWS Glue Schema Registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#schema_reference GlueTable#schema_reference}
	SchemaReference *GlueTableTableInputStorageDescriptorSchemaReference `field:"optional" json:"schemaReference" yaml:"schemaReference"`
	// The serialization/deserialization (SerDe) information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#serde_info GlueTable#serde_info}
	SerdeInfo *GlueTableTableInputStorageDescriptorSerdeInfo `field:"optional" json:"serdeInfo" yaml:"serdeInfo"`
	// The information about values that appear frequently in a column (skewed values).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#skewed_info GlueTable#skewed_info}
	SkewedInfo *GlueTableTableInputStorageDescriptorSkewedInfo `field:"optional" json:"skewedInfo" yaml:"skewedInfo"`
	// A list specifying the sort order of each bucket in the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#sort_columns GlueTable#sort_columns}
	SortColumns interface{} `field:"optional" json:"sortColumns" yaml:"sortColumns"`
	// True if the table data is stored in subdirectories, or False if not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/glue_table#stored_as_sub_directories GlueTable#stored_as_sub_directories}
	StoredAsSubDirectories interface{} `field:"optional" json:"storedAsSubDirectories" yaml:"storedAsSubDirectories"`
}

