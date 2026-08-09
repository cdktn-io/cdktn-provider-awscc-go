// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewjob


type DatabrewJobOutputs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/databrew_job#compression_format DatabrewJob#compression_format}.
	CompressionFormat *string `field:"optional" json:"compressionFormat" yaml:"compressionFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/databrew_job#format DatabrewJob#format}.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Format options for job Output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/databrew_job#format_options DatabrewJob#format_options}
	FormatOptions *DatabrewJobOutputsFormatOptions `field:"optional" json:"formatOptions" yaml:"formatOptions"`
	// S3 Output location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/databrew_job#location DatabrewJob#location}
	Location *DatabrewJobOutputsLocation `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/databrew_job#max_output_files DatabrewJob#max_output_files}.
	MaxOutputFiles *float64 `field:"optional" json:"maxOutputFiles" yaml:"maxOutputFiles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/databrew_job#overwrite DatabrewJob#overwrite}.
	Overwrite interface{} `field:"optional" json:"overwrite" yaml:"overwrite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/databrew_job#partition_columns DatabrewJob#partition_columns}.
	PartitionColumns *[]*string `field:"optional" json:"partitionColumns" yaml:"partitionColumns"`
}

