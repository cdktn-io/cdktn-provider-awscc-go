// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfiguration struct {
	// Cursor-based pagination configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#cursor_configuration GlueConnectionType#cursor_configuration}
	CursorConfiguration *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationCursorConfiguration `field:"optional" json:"cursorConfiguration" yaml:"cursorConfiguration"`
	// Offset-based pagination configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#offset_configuration GlueConnectionType#offset_configuration}
	OffsetConfiguration *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfiguration `field:"optional" json:"offsetConfiguration" yaml:"offsetConfiguration"`
}

