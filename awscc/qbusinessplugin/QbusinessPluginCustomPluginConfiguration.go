// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinessplugin


type QbusinessPluginCustomPluginConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/qbusiness_plugin#api_schema QbusinessPlugin#api_schema}.
	ApiSchema *QbusinessPluginCustomPluginConfigurationApiSchema `field:"optional" json:"apiSchema" yaml:"apiSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/qbusiness_plugin#api_schema_type QbusinessPlugin#api_schema_type}.
	ApiSchemaType *string `field:"optional" json:"apiSchemaType" yaml:"apiSchemaType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/qbusiness_plugin#description QbusinessPlugin#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

