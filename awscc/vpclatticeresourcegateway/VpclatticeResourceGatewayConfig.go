// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vpclatticeresourcegateway

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VpclatticeResourceGatewayConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/vpclattice_resource_gateway#name VpclatticeResourceGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of one or more subnets in which to create an endpoint network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/vpclattice_resource_gateway#subnet_ids VpclatticeResourceGateway#subnet_ids}
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/vpclattice_resource_gateway#vpc_identifier VpclatticeResourceGateway#vpc_identifier}.
	VpcIdentifier *string `field:"required" json:"vpcIdentifier" yaml:"vpcIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/vpclattice_resource_gateway#ip_address_type VpclatticeResourceGateway#ip_address_type}.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// The number of IPv4 addresses to allocate per ENI for the resource gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/vpclattice_resource_gateway#ipv_4_addresses_per_eni VpclatticeResourceGateway#ipv_4_addresses_per_eni}
	Ipv4AddressesPerEni *float64 `field:"optional" json:"ipv4AddressesPerEni" yaml:"ipv4AddressesPerEni"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/vpclattice_resource_gateway#resource_config_dns_resolution VpclatticeResourceGateway#resource_config_dns_resolution}.
	ResourceConfigDnsResolution *string `field:"optional" json:"resourceConfigDnsResolution" yaml:"resourceConfigDnsResolution"`
	// The ID of one or more security groups to associate with the endpoint network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/vpclattice_resource_gateway#security_group_ids VpclatticeResourceGateway#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/vpclattice_resource_gateway#tags VpclatticeResourceGateway#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

