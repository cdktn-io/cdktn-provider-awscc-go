// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationdataflowendpointgroupv2


type GroundstationDataflowEndpointGroupV2EndpointsDownlinkAwsGroundStationAgentEndpointDataflowDetailsAgentConnectionDetails struct {
	// Socket address of an uplink or downlink agent endpoint with a port range and an optional mtu.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/groundstation_dataflow_endpoint_group_v2#agent_ip_and_port_address GroundstationDataflowEndpointGroupV2#agent_ip_and_port_address}
	AgentIpAndPortAddress *GroundstationDataflowEndpointGroupV2EndpointsDownlinkAwsGroundStationAgentEndpointDataflowDetailsAgentConnectionDetailsAgentIpAndPortAddress `field:"optional" json:"agentIpAndPortAddress" yaml:"agentIpAndPortAddress"`
	// Socket address of an uplink or downlink agent endpoint with an optional mtu.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/groundstation_dataflow_endpoint_group_v2#egress_address_and_port GroundstationDataflowEndpointGroupV2#egress_address_and_port}
	EgressAddressAndPort *GroundstationDataflowEndpointGroupV2EndpointsDownlinkAwsGroundStationAgentEndpointDataflowDetailsAgentConnectionDetailsEgressAddressAndPort `field:"optional" json:"egressAddressAndPort" yaml:"egressAddressAndPort"`
}

