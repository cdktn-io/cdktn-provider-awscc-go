// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskchannel


type MskChannelIcebergDestinationConfigurationSchemaEvolution struct {
	// Whether schema evolution is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_channel#enable_schema_evolution MskChannel#enable_schema_evolution}
	EnableSchemaEvolution interface{} `field:"optional" json:"enableSchemaEvolution" yaml:"enableSchemaEvolution"`
}

