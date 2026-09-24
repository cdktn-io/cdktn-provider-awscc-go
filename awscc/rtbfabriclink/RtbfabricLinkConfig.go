// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabriclink

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RtbfabricLinkConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/rtbfabric_link#gateway_id RtbfabricLink#gateway_id}.
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/rtbfabric_link#link_log_settings RtbfabricLink#link_log_settings}.
	LinkLogSettings *RtbfabricLinkLinkLogSettings `field:"required" json:"linkLogSettings" yaml:"linkLogSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/rtbfabric_link#peer_gateway_id RtbfabricLink#peer_gateway_id}.
	PeerGatewayId *string `field:"required" json:"peerGatewayId" yaml:"peerGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/rtbfabric_link#http_responder_allowed RtbfabricLink#http_responder_allowed}.
	HttpResponderAllowed interface{} `field:"optional" json:"httpResponderAllowed" yaml:"httpResponderAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/rtbfabric_link#link_attributes RtbfabricLink#link_attributes}.
	LinkAttributes *RtbfabricLinkLinkAttributes `field:"optional" json:"linkAttributes" yaml:"linkAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/rtbfabric_link#module_configuration_list RtbfabricLink#module_configuration_list}.
	ModuleConfigurationList interface{} `field:"optional" json:"moduleConfigurationList" yaml:"moduleConfigurationList"`
	// Tags to assign to the Link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/rtbfabric_link#tags RtbfabricLink#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

