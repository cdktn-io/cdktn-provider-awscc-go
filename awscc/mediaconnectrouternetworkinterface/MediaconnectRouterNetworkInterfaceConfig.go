// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouternetworkinterface

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectRouterNetworkInterfaceConfig struct {
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
	// The configuration settings for a router network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_router_network_interface#configuration MediaconnectRouterNetworkInterface#configuration}
	Configuration *MediaconnectRouterNetworkInterfaceConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// The name of the router network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_router_network_interface#name MediaconnectRouterNetworkInterface#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The AWS Region for the router network interface. Defaults to the current region if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_router_network_interface#region_name MediaconnectRouterNetworkInterface#region_name}
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// Key-value pairs that can be used to tag and organize this router network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_router_network_interface#tags MediaconnectRouterNetworkInterface#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

