// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueschema


type GlueSchemaRegistry struct {
	// Amazon Resource Name for the Registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/glue_schema#arn GlueSchema#arn}
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Name of the registry in which the schema will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/glue_schema#name GlueSchema#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

