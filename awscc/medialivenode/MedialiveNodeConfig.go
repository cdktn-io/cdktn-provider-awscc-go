// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialivenode

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MedialiveNodeConfig struct {
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
	// The ID of the Cluster that the Node belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/medialive_node#cluster_id MedialiveNode#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The user-specified name of the Node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/medialive_node#name MedialiveNode#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of interface mappings for the Node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/medialive_node#node_interface_mappings MedialiveNode#node_interface_mappings}
	NodeInterfaceMappings interface{} `field:"optional" json:"nodeInterfaceMappings" yaml:"nodeInterfaceMappings"`
	// The role of the Node in the Cluster.
	//
	// ACTIVE means the Node is available for encoding. BACKUP means the Node is a redundant Node and might get used if an ACTIVE Node fails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/medialive_node#role MedialiveNode#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// An array of SDI source mappings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/medialive_node#sdi_source_mappings MedialiveNode#sdi_source_mappings}
	SdiSourceMappings interface{} `field:"optional" json:"sdiSourceMappings" yaml:"sdiSourceMappings"`
	// A collection of key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/medialive_node#tags MedialiveNode#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

