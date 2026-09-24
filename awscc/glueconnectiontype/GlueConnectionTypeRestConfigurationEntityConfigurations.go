// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeRestConfigurationEntityConfigurations struct {
	// The schema definition for this entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#schema GlueConnectionType#schema}
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
	// Configuration that defines how to make requests to endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#source_configuration GlueConnectionType#source_configuration}
	SourceConfiguration *GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfiguration `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
}

