// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeRestConfigurationGlobalSourceConfiguration struct {
	// Configuration that defines how filter predicates are applied to REST API requests, supporting both query parameter and filter string strategies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#filter_configuration GlueConnectionType#filter_configuration}
	FilterConfiguration *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfiguration `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
	// Configuration for handling paginated responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#pagination_configuration GlueConnectionType#pagination_configuration}
	PaginationConfiguration *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfiguration `field:"optional" json:"paginationConfiguration" yaml:"paginationConfiguration"`
	// The HTTP method to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#request_method GlueConnectionType#request_method}
	RequestMethod *string `field:"optional" json:"requestMethod" yaml:"requestMethod"`
	// Request parameters configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#request_parameters GlueConnectionType#request_parameters}
	RequestParameters interface{} `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// The URL path for the REST endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#request_path GlueConnectionType#request_path}
	RequestPath *string `field:"optional" json:"requestPath" yaml:"requestPath"`
	// Configuration for parsing JSON responses from REST API calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#response_configuration GlueConnectionType#response_configuration}
	ResponseConfiguration *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationResponseConfiguration `field:"optional" json:"responseConfiguration" yaml:"responseConfiguration"`
}

