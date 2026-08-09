// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logsquerydefinition

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LogsQueryDefinitionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A name for the saved query definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/logs_query_definition#name LogsQueryDefinition#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The query string to use for this definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/logs_query_definition#query_string LogsQueryDefinition#query_string}
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// Optionally define specific log groups as part of your query definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/logs_query_definition#log_group_names LogsQueryDefinition#log_group_names}
	LogGroupNames *[]*string `field:"optional" json:"logGroupNames" yaml:"logGroupNames"`
	// Use this parameter to include specific query parameters as part of your query definition.
	//
	// Query parameters are supported only for Logs Insights QL queries. Query parameters allow you to use placeholder variables in your query string that are substituted with values at execution time. Use the {{parameterName}} syntax in your query string to reference a parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/logs_query_definition#parameters LogsQueryDefinition#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Query language of the query string. Possible values are CWLI, SQL, PPL, with CWLI being the default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/logs_query_definition#query_language LogsQueryDefinition#query_language}
	QueryLanguage *string `field:"optional" json:"queryLanguage" yaml:"queryLanguage"`
}

