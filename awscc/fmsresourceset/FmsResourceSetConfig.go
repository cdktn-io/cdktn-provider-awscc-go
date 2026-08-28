// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fmsresourceset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FmsResourceSetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/fms_resource_set#name FmsResourceSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/fms_resource_set#resource_type_list FmsResourceSet#resource_type_list}.
	ResourceTypeList *[]*string `field:"required" json:"resourceTypeList" yaml:"resourceTypeList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/fms_resource_set#description FmsResourceSet#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/fms_resource_set#resources FmsResourceSet#resources}.
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/fms_resource_set#tags FmsResourceSet#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

