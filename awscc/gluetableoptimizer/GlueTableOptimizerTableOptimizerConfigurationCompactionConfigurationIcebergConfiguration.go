// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetableoptimizer


type GlueTableOptimizerTableOptimizerConfigurationCompactionConfigurationIcebergConfiguration struct {
	// The minimum number of deletes in a data file to make it eligible for compaction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_table_optimizer#delete_file_threshold GlueTableOptimizer#delete_file_threshold}
	DeleteFileThreshold *float64 `field:"optional" json:"deleteFileThreshold" yaml:"deleteFileThreshold"`
	// The minimum number of input files before compaction is triggered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_table_optimizer#min_input_files GlueTableOptimizer#min_input_files}
	MinInputFiles *float64 `field:"optional" json:"minInputFiles" yaml:"minInputFiles"`
	// The compaction strategy to use. Valid values are binpack, sort, and z-order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_table_optimizer#strategy GlueTableOptimizer#strategy}
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

