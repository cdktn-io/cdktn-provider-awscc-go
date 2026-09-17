// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeConnectionProperties struct {
	// Key-value pairs of additional request parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#additional_request_parameters GlueConnectionType#additional_request_parameters}
	AdditionalRequestParameters interface{} `field:"optional" json:"additionalRequestParameters" yaml:"additionalRequestParameters"`
	// Defines a property configuration for connection types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#url GlueConnectionType#url}
	Url *GlueConnectionTypeConnectionPropertiesUrl `field:"optional" json:"url" yaml:"url"`
}

