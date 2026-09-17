// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2networkinterface


type Ec2NetworkInterfaceConnectionTrackingSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_network_interface#tcp_established_timeout Ec2NetworkInterface#tcp_established_timeout}.
	TcpEstablishedTimeout *float64 `field:"optional" json:"tcpEstablishedTimeout" yaml:"tcpEstablishedTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_network_interface#udp_stream_timeout Ec2NetworkInterface#udp_stream_timeout}.
	UdpStreamTimeout *float64 `field:"optional" json:"udpStreamTimeout" yaml:"udpStreamTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_network_interface#udp_timeout Ec2NetworkInterface#udp_timeout}.
	UdpTimeout *float64 `field:"optional" json:"udpTimeout" yaml:"udpTimeout"`
}

