// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagerglobalnetwork

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkmanagerGlobalNetworkConfig struct {
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
	// The date and time that the global network was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/networkmanager_global_network#created_at NetworkmanagerGlobalNetwork#created_at}
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// The description of the global network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/networkmanager_global_network#description NetworkmanagerGlobalNetwork#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The state of the global network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/networkmanager_global_network#state NetworkmanagerGlobalNetwork#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// The tags for the global network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/networkmanager_global_network#tags NetworkmanagerGlobalNetwork#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

