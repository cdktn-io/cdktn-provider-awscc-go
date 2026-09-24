// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaFilterOverrides struct {
	// Configuration that defines how BETWEEN range filter operations are translated into REST API request parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#between_configuration GlueConnectionType#between_configuration}
	BetweenConfiguration *GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaFilterOverridesBetweenConfiguration `field:"optional" json:"betweenConfiguration" yaml:"betweenConfiguration"`
	// The date and time format for filter expressions on this field, overriding the global DateTimeFormat.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#date_time_format GlueConnectionType#date_time_format}
	DateTimeFormat *string `field:"optional" json:"dateTimeFormat" yaml:"dateTimeFormat"`
	// An override for the field name to use in filter expressions, if different from the schema field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#field_name GlueConnectionType#field_name}
	FieldName *string `field:"optional" json:"fieldName" yaml:"fieldName"`
	// A map of logical filter operators to their field-specific API representations, overriding the global operator mappings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#operator_mappings GlueConnectionType#operator_mappings}
	OperatorMappings *map[string]*string `field:"optional" json:"operatorMappings" yaml:"operatorMappings"`
}

