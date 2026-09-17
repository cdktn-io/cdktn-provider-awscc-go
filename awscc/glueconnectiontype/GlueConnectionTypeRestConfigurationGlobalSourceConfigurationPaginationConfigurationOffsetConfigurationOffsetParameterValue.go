// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValue struct {
	// A JSON path expression to extract a value from response body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#content_path GlueConnectionType#content_path}
	ContentPath *string `field:"optional" json:"contentPath" yaml:"contentPath"`
	// The name of an HTTP response header from which to extract the value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#header_key GlueConnectionType#header_key}
	HeaderKey *string `field:"optional" json:"headerKey" yaml:"headerKey"`
}

