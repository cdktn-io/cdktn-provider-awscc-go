// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagercorenetworkprefixlistassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkmanagerCoreNetworkPrefixListAssociationConfig struct {
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
	// The ID of the core network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkmanager_core_network_prefix_list_association#core_network_id NetworkmanagerCoreNetworkPrefixListAssociation#core_network_id}
	CoreNetworkId *string `field:"required" json:"coreNetworkId" yaml:"coreNetworkId"`
	// The alias of the prefix list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkmanager_core_network_prefix_list_association#prefix_list_alias NetworkmanagerCoreNetworkPrefixListAssociation#prefix_list_alias}
	PrefixListAlias *string `field:"required" json:"prefixListAlias" yaml:"prefixListAlias"`
	// The Amazon Resource Name (ARN) of the prefix list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkmanager_core_network_prefix_list_association#prefix_list_arn NetworkmanagerCoreNetworkPrefixListAssociation#prefix_list_arn}
	PrefixListArn *string `field:"required" json:"prefixListArn" yaml:"prefixListArn"`
}

