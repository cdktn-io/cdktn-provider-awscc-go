// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2transitgatewayattachmentnatconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2TransitGatewayAttachmentNatConfigurationConfig struct {
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
	// The complete set of NAT rules for the attachment.
	//
	// Updating this property replaces the entire rule set: rules omitted from the new value are removed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_transit_gateway_attachment_nat_configuration#nat_gateway_rules Ec2TransitGatewayAttachmentNatConfiguration#nat_gateway_rules}
	NatGatewayRules interface{} `field:"required" json:"natGatewayRules" yaml:"natGatewayRules"`
	// Whether NAT gateway support is enabled on the attachment.
	//
	// Set to disable to stage NAT rules without activating them; the rules are retained and become active when this is later set to enable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_transit_gateway_attachment_nat_configuration#nat_gateway_support Ec2TransitGatewayAttachmentNatConfiguration#nat_gateway_support}
	NatGatewaySupport *string `field:"required" json:"natGatewaySupport" yaml:"natGatewaySupport"`
	// The ID of the transit gateway attachment to configure NAT on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_transit_gateway_attachment_nat_configuration#transit_gateway_attachment_id Ec2TransitGatewayAttachmentNatConfiguration#transit_gateway_attachment_id}
	TransitGatewayAttachmentId *string `field:"required" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
	// Whether routes for the NAT CIDRs are installed into the transit gateway route tables associated with the attachment.
	//
	// Whether this is configurable depends on the attachment; where the value is forced, supplying the opposite value is rejected. Omitting this property leaves the setting to the service: on a new configuration the service applies the attachment's default, but on an attachment where the value is configurable it keeps the value already in effect rather than reverting to that default. Removing this property from a template, or deleting this resource, therefore leaves the setting at its last value rather than restoring the default, and a resource later created on the same attachment without this property inherits that value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_transit_gateway_attachment_nat_configuration#nat_system_route_support Ec2TransitGatewayAttachmentNatConfiguration#nat_system_route_support}
	NatSystemRouteSupport *string `field:"optional" json:"natSystemRouteSupport" yaml:"natSystemRouteSupport"`
}

