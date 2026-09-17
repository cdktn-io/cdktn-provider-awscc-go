// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationFilterConfigurationFilterStringConfiguration struct {
	// The query parameter name used to send the constructed filter expression string in API requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#query_parameter_name GlueConnectionType#query_parameter_name}
	QueryParameterName *string `field:"optional" json:"queryParameterName" yaml:"queryParameterName"`
	// The character used to quote values when QuoteStringValues is true. Defaults to double quotes if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#quote_character GlueConnectionType#quote_character}
	QuoteCharacter *string `field:"optional" json:"quoteCharacter" yaml:"quoteCharacter"`
	// Indicates whether string and date values should be wrapped with a quote character in the filter expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#quote_string_values GlueConnectionType#quote_string_values}
	QuoteStringValues interface{} `field:"optional" json:"quoteStringValues" yaml:"quoteStringValues"`
}

