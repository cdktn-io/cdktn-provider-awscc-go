// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectdestinationendpoint


type MediaconnectDestinationEndpointConfigurationStandardDestinationNetworkingConfigurationVpc struct {
	// The IDs of the security groups to associate with the destination endpoint.
	//
	// You can specify from 1 to 5 security groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#security_group_ids MediaconnectDestinationEndpoint#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnet to use for the destination endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#subnet_id MediaconnectDestinationEndpoint#subnet_id}
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

