// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetableoptimizer


type GlueTableOptimizerTableOptimizerConfigurationCompactionConfiguration struct {
	// The configuration for an Iceberg compaction optimizer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_table_optimizer#iceberg_configuration GlueTableOptimizer#iceberg_configuration}
	IcebergConfiguration *GlueTableOptimizerTableOptimizerConfigurationCompactionConfigurationIcebergConfiguration `field:"optional" json:"icebergConfiguration" yaml:"icebergConfiguration"`
}

