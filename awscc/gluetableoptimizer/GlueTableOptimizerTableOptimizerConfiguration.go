// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetableoptimizer


type GlueTableOptimizerTableOptimizerConfiguration struct {
	// Whether the table optimization is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_table_optimizer#enabled GlueTableOptimizer#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// A role passed by the caller which gives the service permission to update the resources associated with the optimizer on the caller's behalf.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_table_optimizer#role_arn GlueTableOptimizer#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The configuration for a compaction optimizer.
	//
	// This configuration defines how data files in your table will be compacted to improve query performance and reduce storage costs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_table_optimizer#compaction_configuration GlueTableOptimizer#compaction_configuration}
	CompactionConfiguration *GlueTableOptimizerTableOptimizerConfigurationCompactionConfiguration `field:"optional" json:"compactionConfiguration" yaml:"compactionConfiguration"`
	// OrphanFileDeletionConfiguration is a property that can be included within the TableOptimizer resource.
	//
	// It controls the automatic deletion of orphaned files - files that are not tracked by the table metadata, and older than the configured age limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_table_optimizer#orphan_file_deletion_configuration GlueTableOptimizer#orphan_file_deletion_configuration}
	OrphanFileDeletionConfiguration *GlueTableOptimizerTableOptimizerConfigurationOrphanFileDeletionConfiguration `field:"optional" json:"orphanFileDeletionConfiguration" yaml:"orphanFileDeletionConfiguration"`
	// The configuration for a snapshot retention optimizer for Apache Iceberg tables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_table_optimizer#retention_configuration GlueTableOptimizer#retention_configuration}
	RetentionConfiguration *GlueTableOptimizerTableOptimizerConfigurationRetentionConfiguration `field:"optional" json:"retentionConfiguration" yaml:"retentionConfiguration"`
	// An object that describes the VPC configuration for a table optimizer.
	//
	// This configuration is necessary to perform optimization on tables that are in a customer VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_table_optimizer#vpc_configuration GlueTableOptimizer#vpc_configuration}
	VpcConfiguration *GlueTableOptimizerTableOptimizerConfigurationVpcConfiguration `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

