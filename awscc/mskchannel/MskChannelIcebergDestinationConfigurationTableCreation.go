// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskchannel


type MskChannelIcebergDestinationConfigurationTableCreation struct {
	// Whether table creation is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/msk_channel#enable_table_creation MskChannel#enable_table_creation}
	EnableTableCreation interface{} `field:"optional" json:"enableTableCreation" yaml:"enableTableCreation"`
}

