// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfiguration struct {
	// Configuration that defines how BETWEEN range filter operations are translated into REST API request parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#between_configuration GlueConnectionType#between_configuration}
	BetweenConfiguration *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfiguration `field:"optional" json:"betweenConfiguration" yaml:"betweenConfiguration"`
	// The global date and time format for filter expressions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#date_time_format GlueConnectionType#date_time_format}
	DateTimeFormat *string `field:"optional" json:"dateTimeFormat" yaml:"dateTimeFormat"`
	// The strategy for applying filters to requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#filter_mode GlueConnectionType#filter_mode}
	FilterMode *string `field:"optional" json:"filterMode" yaml:"filterMode"`
	// Configuration for constructing filter expression strings when using the FILTER_STRING filter mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#filter_string_configuration GlueConnectionType#filter_string_configuration}
	FilterStringConfiguration *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationFilterStringConfiguration `field:"optional" json:"filterStringConfiguration" yaml:"filterStringConfiguration"`
	// A map of logical filter operators to their API-specific string representations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#operator_mappings GlueConnectionType#operator_mappings}
	OperatorMappings *map[string]*string `field:"optional" json:"operatorMappings" yaml:"operatorMappings"`
	// Indicates whether surrounding double quotes should be stripped from filter values before processing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#strip_quotes GlueConnectionType#strip_quotes}
	StripQuotes interface{} `field:"optional" json:"stripQuotes" yaml:"stripQuotes"`
}

