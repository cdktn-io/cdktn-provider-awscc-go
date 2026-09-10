// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package neptunegraphgraphsnapshot

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NeptunegraphGraphSnapshotConfig struct {
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
	// The unique identifier of the Neptune Analytics graph to create the snapshot from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/neptunegraph_graph_snapshot#graph_identifier NeptunegraphGraphSnapshot#graph_identifier}
	GraphIdentifier *string `field:"required" json:"graphIdentifier" yaml:"graphIdentifier"`
	// The snapshot name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/neptunegraph_graph_snapshot#snapshot_name NeptunegraphGraphSnapshot#snapshot_name}
	SnapshotName *string `field:"required" json:"snapshotName" yaml:"snapshotName"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/neptunegraph_graph_snapshot#tags NeptunegraphGraphSnapshot#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

