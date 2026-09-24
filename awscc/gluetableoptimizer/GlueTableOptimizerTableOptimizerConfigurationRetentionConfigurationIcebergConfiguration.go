// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetableoptimizer


type GlueTableOptimizerTableOptimizerConfigurationRetentionConfigurationIcebergConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_table_optimizer#clean_expired_files GlueTableOptimizer#clean_expired_files}.
	CleanExpiredFiles interface{} `field:"optional" json:"cleanExpiredFiles" yaml:"cleanExpiredFiles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_table_optimizer#number_of_snapshots_to_retain GlueTableOptimizer#number_of_snapshots_to_retain}.
	NumberOfSnapshotsToRetain *float64 `field:"optional" json:"numberOfSnapshotsToRetain" yaml:"numberOfSnapshotsToRetain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_table_optimizer#snapshot_retention_period_in_days GlueTableOptimizer#snapshot_retention_period_in_days}.
	SnapshotRetentionPeriodInDays *float64 `field:"optional" json:"snapshotRetentionPeriodInDays" yaml:"snapshotRetentionPeriodInDays"`
}

