// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeRestConfigurationGlobalSourceConfigurationResponseConfiguration struct {
	// JSON path expression for error information location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#error_path GlueConnectionType#error_path}
	ErrorPath *string `field:"optional" json:"errorPath" yaml:"errorPath"`
	// JSON path expression for result data location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#result_path GlueConnectionType#result_path}
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
}

