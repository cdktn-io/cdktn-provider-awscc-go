// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfiguration struct {
	// Parameter extraction configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#limit_parameter GlueConnectionType#limit_parameter}
	LimitParameter *GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationLimitParameter `field:"optional" json:"limitParameter" yaml:"limitParameter"`
	// Parameter extraction configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#offset_parameter GlueConnectionType#offset_parameter}
	OffsetParameter *GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameter `field:"optional" json:"offsetParameter" yaml:"offsetParameter"`
}

