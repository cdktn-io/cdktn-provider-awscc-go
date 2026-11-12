// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectsourceendpoint


type MediaconnectSourceEndpointConfigurationStandardSourceNetworkingConfigurationVpc struct {
	// The IDs of the security groups to associate with the source endpoint.
	//
	// You can specify from 1 to 5 security groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#security_group_ids MediaconnectSourceEndpoint#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnet to use for the source endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#subnet_id MediaconnectSourceEndpoint#subnet_id}
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

