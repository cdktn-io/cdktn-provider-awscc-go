// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package entityresolutionschemamapping

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EntityresolutionSchemaMappingConfig struct {
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
	// The SchemaMapping attributes input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/entityresolution_schema_mapping#mapped_input_fields EntityresolutionSchemaMapping#mapped_input_fields}
	MappedInputFields interface{} `field:"required" json:"mappedInputFields" yaml:"mappedInputFields"`
	// The name of the SchemaMapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/entityresolution_schema_mapping#schema_name EntityresolutionSchemaMapping#schema_name}
	SchemaName *string `field:"required" json:"schemaName" yaml:"schemaName"`
	// The description of the SchemaMapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/entityresolution_schema_mapping#description EntityresolutionSchemaMapping#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/entityresolution_schema_mapping#tags EntityresolutionSchemaMapping#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

