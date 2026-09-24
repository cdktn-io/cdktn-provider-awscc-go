// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fmsprotocolslist

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FmsProtocolsListConfig struct {
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
	// The name of the Firewall Manager protocols list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fms_protocols_list#list_name FmsProtocolsList#list_name}
	ListName *string `field:"required" json:"listName" yaml:"listName"`
	// An array of protocols in the Firewall Manager protocols list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fms_protocols_list#protocols_list FmsProtocolsList#protocols_list}
	ProtocolsList *[]*string `field:"required" json:"protocolsList" yaml:"protocolsList"`
	// An array of key-value pairs to apply to the protocols list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fms_protocols_list#tags FmsProtocolsList#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

