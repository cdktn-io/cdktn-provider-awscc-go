// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentprivateconnection


type DevopsagentPrivateConnectionConnectionConfigurationServiceManaged struct {
	// DNS resolution mode for the resource gateway. Defaults to PUBLIC when not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_private_connection#dns_resolution DevopsagentPrivateConnection#dns_resolution}
	DnsResolution *string `field:"optional" json:"dnsResolution" yaml:"dnsResolution"`
	// IP address or DNS name of the target resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_private_connection#host_address DevopsagentPrivateConnection#host_address}
	HostAddress *string `field:"optional" json:"hostAddress" yaml:"hostAddress"`
	// IP address type of the service-managed Resource Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_private_connection#ip_address_type DevopsagentPrivateConnection#ip_address_type}
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Number of IPv4 addresses in each ENI for the service-managed Resource Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_private_connection#ipv_4_addresses_per_eni DevopsagentPrivateConnection#ipv_4_addresses_per_eni}
	Ipv4AddressesPerEni *float64 `field:"optional" json:"ipv4AddressesPerEni" yaml:"ipv4AddressesPerEni"`
	// TCP port ranges that a consumer can use to access the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_private_connection#port_ranges DevopsagentPrivateConnection#port_ranges}
	PortRanges *[]*string `field:"optional" json:"portRanges" yaml:"portRanges"`
	// Security groups to attach to the service-managed Resource Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_private_connection#security_group_ids DevopsagentPrivateConnection#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Subnets that the service-managed Resource Gateway will span.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_private_connection#subnet_ids DevopsagentPrivateConnection#subnet_ids}
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// VPC to create the service-managed Resource Gateway in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_private_connection#vpc_id DevopsagentPrivateConnection#vpc_id}
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

