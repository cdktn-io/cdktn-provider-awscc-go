// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewaydomainnamev2


type ApigatewayDomainNameV2EndpointConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_domain_name_v2#ip_address_type ApigatewayDomainNameV2#ip_address_type}.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_domain_name_v2#types ApigatewayDomainNameV2#types}.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

