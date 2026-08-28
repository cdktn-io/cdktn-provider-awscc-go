// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2capacitymanagerdataexport

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2CapacityManagerDataExportConfig struct {
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
	// The format of the exported capacity manager data.
	//
	// Choose 'csv' for comma-separated values or 'parquet' for optimized columnar storage format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_capacity_manager_data_export#output_format Ec2CapacityManagerDataExport#output_format}
	OutputFormat *string `field:"required" json:"outputFormat" yaml:"outputFormat"`
	// The name of the Amazon S3 bucket where the capacity manager data export will be stored.
	//
	// The bucket must exist and be accessible by EC2 Capacity Manager service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_capacity_manager_data_export#s3_bucket_name Ec2CapacityManagerDataExport#s3_bucket_name}
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// The schedule for the capacity manager data export.
	//
	// Currently supports hourly exports that provide periodic snapshots of capacity manager data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_capacity_manager_data_export#schedule Ec2CapacityManagerDataExport#schedule}
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// The prefix for the S3 bucket location where exported files will be placed.
	//
	// If not specified, files will be placed in the root of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_capacity_manager_data_export#s3_bucket_prefix Ec2CapacityManagerDataExport#s3_bucket_prefix}
	S3BucketPrefix *string `field:"optional" json:"s3BucketPrefix" yaml:"s3BucketPrefix"`
	// An array of key-value pairs to apply to the capacity manager data export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_capacity_manager_data_export#tags Ec2CapacityManagerDataExport#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

