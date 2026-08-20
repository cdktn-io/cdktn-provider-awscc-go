// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ippoolroutetableassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2IpPoolRouteTableAssociationConfig struct {
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
	// The ID of the public IPv4 pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_ip_pool_route_table_association#public_ipv_4_pool Ec2IpPoolRouteTableAssociation#public_ipv_4_pool}
	PublicIpv4Pool *string `field:"required" json:"publicIpv4Pool" yaml:"publicIpv4Pool"`
	// The ID of the route table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_ip_pool_route_table_association#route_table_id Ec2IpPoolRouteTableAssociation#route_table_id}
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
}

