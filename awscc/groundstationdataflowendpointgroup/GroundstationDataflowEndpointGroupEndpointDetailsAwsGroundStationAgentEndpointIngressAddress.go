// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationdataflowendpointgroup


type GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointIngressAddress struct {
	// Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_dataflow_endpoint_group#mtu GroundstationDataflowEndpointGroup#mtu}
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// A socket address with a port range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_dataflow_endpoint_group#socket_address GroundstationDataflowEndpointGroup#socket_address}
	SocketAddress *GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointIngressAddressSocketAddress `field:"optional" json:"socketAddress" yaml:"socketAddress"`
}

