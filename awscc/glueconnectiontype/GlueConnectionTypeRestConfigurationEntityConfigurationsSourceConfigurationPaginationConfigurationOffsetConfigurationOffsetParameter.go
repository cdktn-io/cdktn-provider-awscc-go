// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameter struct {
	// The default value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#default_value GlueConnectionType#default_value}
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The parameter key name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#key GlueConnectionType#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Specifies where to place the parameter in requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#property_location GlueConnectionType#property_location}
	PropertyLocation *string `field:"optional" json:"propertyLocation" yaml:"propertyLocation"`
	// Defines how to extract values from HTTP responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#value GlueConnectionType#value}
	Value *GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValue `field:"optional" json:"value" yaml:"value"`
}

