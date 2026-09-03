// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationdataflowendpointgroupv2


type GroundstationDataflowEndpointGroupV2EndpointsDownlinkAwsGroundStationAgentEndpointDataflowDetailsAgentConnectionDetailsEgressAddressAndPort struct {
	// Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/groundstation_dataflow_endpoint_group_v2#mtu GroundstationDataflowEndpointGroupV2#mtu}
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/groundstation_dataflow_endpoint_group_v2#socket_address GroundstationDataflowEndpointGroupV2#socket_address}.
	SocketAddress *GroundstationDataflowEndpointGroupV2EndpointsDownlinkAwsGroundStationAgentEndpointDataflowDetailsAgentConnectionDetailsEgressAddressAndPortSocketAddress `field:"optional" json:"socketAddress" yaml:"socketAddress"`
}

