// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scnnamespace

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ScnNamespaceConfig struct {
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
	// The Amazon Web Services Supply Chain instance identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/scn_namespace#instance_id ScnNamespace#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The name of the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/scn_namespace#name ScnNamespace#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/scn_namespace#description ScnNamespace#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags for the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/scn_namespace#tags ScnNamespace#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

