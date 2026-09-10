// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueschema


type GlueSchemaCheckpointVersion struct {
	// Indicates if the latest version needs to be updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_schema#is_latest GlueSchema#is_latest}
	IsLatest interface{} `field:"optional" json:"isLatest" yaml:"isLatest"`
	// Indicates the version number in the schema to update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_schema#version_number GlueSchema#version_number}
	VersionNumber *float64 `field:"optional" json:"versionNumber" yaml:"versionNumber"`
}

