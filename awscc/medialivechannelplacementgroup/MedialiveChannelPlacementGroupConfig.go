// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialivechannelplacementgroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MedialiveChannelPlacementGroupConfig struct {
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
	// The ID of the cluster the node is on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/medialive_channel_placement_group#cluster_id MedialiveChannelPlacementGroup#cluster_id}
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// The name of the channel placement group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/medialive_channel_placement_group#name MedialiveChannelPlacementGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// List of nodes added to the channel placement group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/medialive_channel_placement_group#nodes MedialiveChannelPlacementGroup#nodes}
	Nodes *[]*string `field:"optional" json:"nodes" yaml:"nodes"`
	// A collection of key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/medialive_channel_placement_group#tags MedialiveChannelPlacementGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

