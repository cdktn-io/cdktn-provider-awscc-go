// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagercorenetwork

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkmanagerCoreNetworkConfig struct {
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
	// The ID of the global network that your core network is a part of.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkmanager_core_network#global_network_id NetworkmanagerCoreNetwork#global_network_id}
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The description of core network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkmanager_core_network#description NetworkmanagerCoreNetwork#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Live policy document for the core network, you must provide PolicyDocument in Json Format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkmanager_core_network#policy_document NetworkmanagerCoreNetwork#policy_document}
	PolicyDocument *string `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// The tags for the global network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkmanager_core_network#tags NetworkmanagerCoreNetwork#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

