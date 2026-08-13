// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanageringresspoint


type SesMailManagerIngressPointNetworkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ses_mail_manager_ingress_point#private_network_configuration SesMailManagerIngressPoint#private_network_configuration}.
	PrivateNetworkConfiguration *SesMailManagerIngressPointNetworkConfigurationPrivateNetworkConfiguration `field:"optional" json:"privateNetworkConfiguration" yaml:"privateNetworkConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ses_mail_manager_ingress_point#public_network_configuration SesMailManagerIngressPoint#public_network_configuration}.
	PublicNetworkConfiguration *SesMailManagerIngressPointNetworkConfigurationPublicNetworkConfiguration `field:"optional" json:"publicNetworkConfiguration" yaml:"publicNetworkConfiguration"`
}

