// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationCursorConfiguration struct {
	// Parameter extraction configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#limit_parameter GlueConnectionType#limit_parameter}
	LimitParameter *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationCursorConfigurationLimitParameter `field:"optional" json:"limitParameter" yaml:"limitParameter"`
	// Parameter extraction configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#next_page GlueConnectionType#next_page}
	NextPage *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationCursorConfigurationNextPage `field:"optional" json:"nextPage" yaml:"nextPage"`
}

