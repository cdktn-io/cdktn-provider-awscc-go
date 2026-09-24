// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaFilterOverridesBetweenConfiguration struct {
	// The parameter name used for the upper bound value in a BETWEEN filter operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#high_bound_key GlueConnectionType#high_bound_key}
	HighBoundKey *string `field:"optional" json:"highBoundKey" yaml:"highBoundKey"`
	// The parameter name used for the lower bound value in a BETWEEN filter operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#low_bound_key GlueConnectionType#low_bound_key}
	LowBoundKey *string `field:"optional" json:"lowBoundKey" yaml:"lowBoundKey"`
	// A template string for constructing the BETWEEN filter expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#template GlueConnectionType#template}
	Template *string `field:"optional" json:"template" yaml:"template"`
}

