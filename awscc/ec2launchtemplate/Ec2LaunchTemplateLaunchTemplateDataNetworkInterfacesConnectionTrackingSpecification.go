// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesConnectionTrackingSpecification struct {
	// Timeout (in seconds) for idle TCP connections in an established state.
	//
	// Min: 60 seconds. Max: 432000 seconds (5 days). Default: 432000 seconds. Recommended: Less than 432000 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_launch_template#tcp_established_timeout Ec2LaunchTemplate#tcp_established_timeout}
	TcpEstablishedTimeout *float64 `field:"optional" json:"tcpEstablishedTimeout" yaml:"tcpEstablishedTimeout"`
	// Timeout (in seconds) for idle UDP flows classified as streams which have seen more than one request-response transaction.
	//
	// Min: 60 seconds. Max: 180 seconds (3 minutes). Default: 180 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_launch_template#udp_stream_timeout Ec2LaunchTemplate#udp_stream_timeout}
	UdpStreamTimeout *float64 `field:"optional" json:"udpStreamTimeout" yaml:"udpStreamTimeout"`
	// Timeout (in seconds) for idle UDP flows that have seen traffic only in a single direction or a single request-response transaction.
	//
	// Min: 30 seconds. Max: 60 seconds. Default: 30 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_launch_template#udp_timeout Ec2LaunchTemplate#udp_timeout}
	UdpTimeout *float64 `field:"optional" json:"udpTimeout" yaml:"udpTimeout"`
}

