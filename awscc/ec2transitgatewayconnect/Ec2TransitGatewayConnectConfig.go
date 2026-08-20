// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2transitgatewayconnect

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2TransitGatewayConnectConfig struct {
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
	// The Connect attachment options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_transit_gateway_connect#options Ec2TransitGatewayConnect#options}
	Options *Ec2TransitGatewayConnectOptions `field:"required" json:"options" yaml:"options"`
	// The ID of the attachment from which the Connect attachment was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_transit_gateway_connect#transport_transit_gateway_attachment_id Ec2TransitGatewayConnect#transport_transit_gateway_attachment_id}
	TransportTransitGatewayAttachmentId *string `field:"required" json:"transportTransitGatewayAttachmentId" yaml:"transportTransitGatewayAttachmentId"`
	// The tags for the attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_transit_gateway_connect#tags Ec2TransitGatewayConnect#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

