// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2clientvpntargetnetworkassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2ClientVpnTargetNetworkAssociationConfig struct {
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
	// The ID of the Client VPN endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_target_network_association#client_vpn_endpoint_id Ec2ClientVpnTargetNetworkAssociation#client_vpn_endpoint_id}
	ClientVpnEndpointId *string `field:"required" json:"clientVpnEndpointId" yaml:"clientVpnEndpointId"`
	// The Availability Zone to associate with the Client VPN endpoint. Used for Transit Gateway-based endpoints in place of SubnetId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_target_network_association#availability_zone Ec2ClientVpnTargetNetworkAssociation#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The ID of the Availability Zone to associate with the Client VPN endpoint.
	//
	// Used for Transit Gateway-based endpoints in place of SubnetId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_target_network_association#availability_zone_id Ec2ClientVpnTargetNetworkAssociation#availability_zone_id}
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// The ID of the subnet to associate with the Client VPN endpoint.
	//
	// Required for VPC-based endpoints. For Transit Gateway-based endpoints, use AvailabilityZone or AvailabilityZoneId instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_target_network_association#subnet_id Ec2ClientVpnTargetNetworkAssociation#subnet_id}
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

