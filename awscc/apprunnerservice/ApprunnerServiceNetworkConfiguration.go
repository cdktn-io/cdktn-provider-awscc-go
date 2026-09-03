// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apprunnerservice


type ApprunnerServiceNetworkConfiguration struct {
	// Network egress configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apprunner_service#egress_configuration ApprunnerService#egress_configuration}
	EgressConfiguration *ApprunnerServiceNetworkConfigurationEgressConfiguration `field:"optional" json:"egressConfiguration" yaml:"egressConfiguration"`
	// Network ingress configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apprunner_service#ingress_configuration ApprunnerService#ingress_configuration}
	IngressConfiguration *ApprunnerServiceNetworkConfigurationIngressConfiguration `field:"optional" json:"ingressConfiguration" yaml:"ingressConfiguration"`
	// App Runner service endpoint IP address type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apprunner_service#ip_address_type ApprunnerService#ip_address_type}
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
}

