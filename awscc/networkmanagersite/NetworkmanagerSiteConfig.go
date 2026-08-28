// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagersite

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkmanagerSiteConfig struct {
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
	// The ID of the global network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkmanager_site#global_network_id NetworkmanagerSite#global_network_id}
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The description of the site.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkmanager_site#description NetworkmanagerSite#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The location of the site.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkmanager_site#location NetworkmanagerSite#location}
	Location *NetworkmanagerSiteLocation `field:"optional" json:"location" yaml:"location"`
	// The tags for the site.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkmanager_site#tags NetworkmanagerSite#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

