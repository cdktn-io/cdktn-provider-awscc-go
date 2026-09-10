// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewaydomainname


type ApigatewayDomainNameEndpointConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/apigateway_domain_name#ip_address_type ApigatewayDomainName#ip_address_type}.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/apigateway_domain_name#types ApigatewayDomainName#types}.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

