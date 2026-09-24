// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package directconnectdirectconnectgateway

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DirectconnectDirectConnectGatewayConfig struct {
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
	// The name of the Direct Connect gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/directconnect_direct_connect_gateway#direct_connect_gateway_name DirectconnectDirectConnectGateway#direct_connect_gateway_name}
	DirectConnectGatewayName *string `field:"required" json:"directConnectGatewayName" yaml:"directConnectGatewayName"`
	// The autonomous system number (ASN) for the Amazon side of the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/directconnect_direct_connect_gateway#amazon_side_asn DirectconnectDirectConnectGateway#amazon_side_asn}
	AmazonSideAsn *string `field:"optional" json:"amazonSideAsn" yaml:"amazonSideAsn"`
	// The tags associated with the Direct Connect gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/directconnect_direct_connect_gateway#tags DirectconnectDirectConnectGateway#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

