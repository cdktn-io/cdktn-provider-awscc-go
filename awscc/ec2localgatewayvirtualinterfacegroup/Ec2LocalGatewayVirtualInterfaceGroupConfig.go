// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2localgatewayvirtualinterfacegroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2LocalGatewayVirtualInterfaceGroupConfig struct {
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
	// The ID of the local gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_local_gateway_virtual_interface_group#local_gateway_id Ec2LocalGatewayVirtualInterfaceGroup#local_gateway_id}
	LocalGatewayId *string `field:"required" json:"localGatewayId" yaml:"localGatewayId"`
	// The Autonomous System Number(ASN) for the local Border Gateway Protocol (BGP).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_local_gateway_virtual_interface_group#local_bgp_asn Ec2LocalGatewayVirtualInterfaceGroup#local_bgp_asn}
	LocalBgpAsn *float64 `field:"optional" json:"localBgpAsn" yaml:"localBgpAsn"`
	// The extended 32-bit ASN for the local BGP configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_local_gateway_virtual_interface_group#local_bgp_asn_extended Ec2LocalGatewayVirtualInterfaceGroup#local_bgp_asn_extended}
	LocalBgpAsnExtended *float64 `field:"optional" json:"localBgpAsnExtended" yaml:"localBgpAsnExtended"`
	// The tags assigned to the virtual interface group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_local_gateway_virtual_interface_group#tags Ec2LocalGatewayVirtualInterfaceGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

