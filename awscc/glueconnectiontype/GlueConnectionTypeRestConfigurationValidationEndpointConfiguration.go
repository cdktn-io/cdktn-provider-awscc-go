// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeRestConfigurationValidationEndpointConfiguration struct {
	// The HTTP method to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#request_method GlueConnectionType#request_method}
	RequestMethod *string `field:"optional" json:"requestMethod" yaml:"requestMethod"`
	// The URL path for the REST endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#request_path GlueConnectionType#request_path}
	RequestPath *string `field:"optional" json:"requestPath" yaml:"requestPath"`
}

