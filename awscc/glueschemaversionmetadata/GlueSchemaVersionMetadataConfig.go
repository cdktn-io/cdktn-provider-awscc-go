// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueschemaversionmetadata

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueSchemaVersionMetadataConfig struct {
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
	// Metadata key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_schema_version_metadata#key GlueSchemaVersionMetadata#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Represents the version ID associated with the schema version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_schema_version_metadata#schema_version_id GlueSchemaVersionMetadata#schema_version_id}
	SchemaVersionId *string `field:"required" json:"schemaVersionId" yaml:"schemaVersionId"`
	// Metadata value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_schema_version_metadata#value GlueSchemaVersionMetadata#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

