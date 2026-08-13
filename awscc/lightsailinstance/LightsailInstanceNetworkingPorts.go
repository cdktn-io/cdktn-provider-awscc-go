// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsailinstance


type LightsailInstanceNetworkingPorts struct {
	// Access Direction for Protocol of the Instance(inbound/outbound).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lightsail_instance#access_direction LightsailInstance#access_direction}
	AccessDirection *string `field:"optional" json:"accessDirection" yaml:"accessDirection"`
	// Access From Protocol of the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lightsail_instance#access_from LightsailInstance#access_from}
	AccessFrom *string `field:"optional" json:"accessFrom" yaml:"accessFrom"`
	// Access Type Protocol of the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lightsail_instance#access_type LightsailInstance#access_type}
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
	// cidr List Aliases.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lightsail_instance#cidr_list_aliases LightsailInstance#cidr_list_aliases}
	CidrListAliases *[]*string `field:"optional" json:"cidrListAliases" yaml:"cidrListAliases"`
	// cidrs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lightsail_instance#cidrs LightsailInstance#cidrs}
	Cidrs *[]*string `field:"optional" json:"cidrs" yaml:"cidrs"`
	// CommonName for Protocol of the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lightsail_instance#common_name LightsailInstance#common_name}
	CommonName *string `field:"optional" json:"commonName" yaml:"commonName"`
	// From Port of the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lightsail_instance#from_port LightsailInstance#from_port}
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// IPv6 Cidrs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lightsail_instance#ipv_6_cidrs LightsailInstance#ipv_6_cidrs}
	Ipv6Cidrs *[]*string `field:"optional" json:"ipv6Cidrs" yaml:"ipv6Cidrs"`
	// Port Protocol of the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lightsail_instance#protocol LightsailInstance#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// To Port of the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lightsail_instance#to_port LightsailInstance#to_port}
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

