// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetableoptimizer


type GlueTableOptimizerTableOptimizerConfigurationOrphanFileDeletionConfiguration struct {
	// The IcebergConfiguration property helps optimize your Iceberg tables in AWS Glue by allowing you to specify format-specific settings that control how data is stored, compressed, and managed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_table_optimizer#iceberg_configuration GlueTableOptimizer#iceberg_configuration}
	IcebergConfiguration *GlueTableOptimizerTableOptimizerConfigurationOrphanFileDeletionConfigurationIcebergConfiguration `field:"optional" json:"icebergConfiguration" yaml:"icebergConfiguration"`
}

