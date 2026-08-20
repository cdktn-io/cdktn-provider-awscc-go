// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetableoptimizer


type GlueTableOptimizerTableOptimizerConfigurationRetentionConfiguration struct {
	// The configuration for an Iceberg snapshot retention optimizer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/glue_table_optimizer#iceberg_configuration GlueTableOptimizer#iceberg_configuration}
	IcebergConfiguration *GlueTableOptimizerTableOptimizerConfigurationRetentionConfigurationIcebergConfiguration `field:"optional" json:"icebergConfiguration" yaml:"icebergConfiguration"`
}

