// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayrestapi


type ApigatewayRestApiEndpointConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#ip_address_type ApigatewayRestApi#ip_address_type}.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#types ApigatewayRestApi#types}.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api#vpc_endpoint_ids ApigatewayRestApi#vpc_endpoint_ids}.
	VpcEndpointIds *[]*string `field:"optional" json:"vpcEndpointIds" yaml:"vpcEndpointIds"`
}

