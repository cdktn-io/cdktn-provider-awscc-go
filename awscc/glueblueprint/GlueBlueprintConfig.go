// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueblueprint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueBlueprintConfig struct {
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
	// Specifies a path in Amazon S3 where the blueprint is published.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_blueprint#blueprint_location GlueBlueprint#blueprint_location}
	BlueprintLocation *string `field:"required" json:"blueprintLocation" yaml:"blueprintLocation"`
	// The name of the blueprint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_blueprint#name GlueBlueprint#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the blueprint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_blueprint#description GlueBlueprint#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags to be applied to this blueprint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_blueprint#tags GlueBlueprint#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

