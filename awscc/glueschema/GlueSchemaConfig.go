// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueschema

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueSchemaConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Compatibility setting for the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_schema#compatibility GlueSchema#compatibility}
	Compatibility *string `field:"required" json:"compatibility" yaml:"compatibility"`
	// Data format name to use for the schema. Accepted values: 'AVRO', 'JSON', 'PROTOBUF'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_schema#data_format GlueSchema#data_format}
	DataFormat *string `field:"required" json:"dataFormat" yaml:"dataFormat"`
	// Name of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_schema#name GlueSchema#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify checkpoint version for update. This is only required to update the Compatibility.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_schema#checkpoint_version GlueSchema#checkpoint_version}
	CheckpointVersion *GlueSchemaCheckpointVersion `field:"optional" json:"checkpointVersion" yaml:"checkpointVersion"`
	// A description of the schema. If description is not provided, there will not be any default value for this.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_schema#description GlueSchema#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Identifier for the registry which the schema is part of.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_schema#registry GlueSchema#registry}
	Registry *GlueSchemaRegistry `field:"optional" json:"registry" yaml:"registry"`
	// Definition for the initial schema version in plain-text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_schema#schema_definition GlueSchema#schema_definition}
	SchemaDefinition *string `field:"optional" json:"schemaDefinition" yaml:"schemaDefinition"`
	// List of tags to tag the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_schema#tags GlueSchema#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

