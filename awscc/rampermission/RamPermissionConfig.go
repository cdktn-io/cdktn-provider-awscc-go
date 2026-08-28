// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rampermission

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RamPermissionConfig struct {
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
	// The name of the permission.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ram_permission#name RamPermission#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policy template for the permission.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ram_permission#policy_template RamPermission#policy_template}
	PolicyTemplate *string `field:"required" json:"policyTemplate" yaml:"policyTemplate"`
	// The resource type this permission can be used with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ram_permission#resource_type RamPermission#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ram_permission#tags RamPermission#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

