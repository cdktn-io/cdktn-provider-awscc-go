// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2instance


type Ec2InstanceMetadataOptions struct {
	// Enables or disables the HTTP metadata endpoint on your instances.
	//
	// If you specify a value of disabled, you cannot access your instance metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_instance#http_endpoint Ec2Instance#http_endpoint}
	HttpEndpoint *string `field:"optional" json:"httpEndpoint" yaml:"httpEndpoint"`
	// Enables or disables the IPv6 endpoint for the instance metadata service.
	//
	// To use this option, the instance must be a Nitro-based instance launched in a subnet that supports IPv6.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_instance#http_protocol_ipv_6 Ec2Instance#http_protocol_ipv_6}
	HttpProtocolIpv6 *string `field:"optional" json:"httpProtocolIpv6" yaml:"httpProtocolIpv6"`
	// The number of network hops that the metadata token can travel. Maximum is 64.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_instance#http_put_response_hop_limit Ec2Instance#http_put_response_hop_limit}
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// Indicates whether IMDSv2 is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_instance#http_tokens Ec2Instance#http_tokens}
	HttpTokens *string `field:"optional" json:"httpTokens" yaml:"httpTokens"`
	// Indicates whether tags from the instance are propagated to the EBS volumes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_instance#instance_metadata_tags Ec2Instance#instance_metadata_tags}
	InstanceMetadataTags *string `field:"optional" json:"instanceMetadataTags" yaml:"instanceMetadataTags"`
}

