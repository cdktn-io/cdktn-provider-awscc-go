// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinessplugin


type QbusinessPluginCustomPluginConfigurationApiSchemaS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/qbusiness_plugin#bucket QbusinessPlugin#bucket}.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/qbusiness_plugin#key QbusinessPlugin#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

