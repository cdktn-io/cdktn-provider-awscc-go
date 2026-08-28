// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabricoutboundexternallink

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RtbfabricOutboundExternalLinkConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_outbound_external_link#gateway_id RtbfabricOutboundExternalLink#gateway_id}.
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_outbound_external_link#link_log_settings RtbfabricOutboundExternalLink#link_log_settings}.
	LinkLogSettings *RtbfabricOutboundExternalLinkLinkLogSettings `field:"required" json:"linkLogSettings" yaml:"linkLogSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_outbound_external_link#public_endpoint RtbfabricOutboundExternalLink#public_endpoint}.
	PublicEndpoint *string `field:"required" json:"publicEndpoint" yaml:"publicEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_outbound_external_link#link_attributes RtbfabricOutboundExternalLink#link_attributes}.
	LinkAttributes *RtbfabricOutboundExternalLinkLinkAttributes `field:"optional" json:"linkAttributes" yaml:"linkAttributes"`
	// Tags to assign to the Link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_outbound_external_link#tags RtbfabricOutboundExternalLink#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

