// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamstack


type AppstreamStackAccessEndpoints struct {
	// The type of interface endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appstream_stack#endpoint_type AppstreamStack#endpoint_type}
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
	// The identifier (ID) of the VPC in which the interface endpoint is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appstream_stack#vpce_id AppstreamStack#vpce_id}
	VpceId *string `field:"optional" json:"vpceId" yaml:"vpceId"`
}

