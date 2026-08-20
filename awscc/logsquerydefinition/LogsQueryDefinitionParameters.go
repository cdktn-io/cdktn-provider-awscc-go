// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logsquerydefinition


type LogsQueryDefinitionParameters struct {
	// The default value to use for this query parameter if no value is supplied at execution time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_query_definition#default_value LogsQueryDefinition#default_value}
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// A description of the query parameter that explains its purpose or expected values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_query_definition#description LogsQueryDefinition#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the query parameter.
	//
	// A query parameter name must start with a letter or underscore, and contain only letters, digits, and underscores.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_query_definition#name LogsQueryDefinition#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

