// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vpclatticeresourceconfiguration


type VpclatticeResourceConfigurationResourceConfigurationDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/vpclattice_resource_configuration#arn_resource VpclatticeResourceConfiguration#arn_resource}.
	ArnResource *string `field:"optional" json:"arnResource" yaml:"arnResource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/vpclattice_resource_configuration#cidr_resource VpclatticeResourceConfiguration#cidr_resource}.
	CidrResource *VpclatticeResourceConfigurationResourceConfigurationDefinitionCidrResource `field:"optional" json:"cidrResource" yaml:"cidrResource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/vpclattice_resource_configuration#dns_resource VpclatticeResourceConfiguration#dns_resource}.
	DnsResource *VpclatticeResourceConfigurationResourceConfigurationDefinitionDnsResource `field:"optional" json:"dnsResource" yaml:"dnsResource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/vpclattice_resource_configuration#ip_resource VpclatticeResourceConfiguration#ip_resource}.
	IpResource *string `field:"optional" json:"ipResource" yaml:"ipResource"`
}

