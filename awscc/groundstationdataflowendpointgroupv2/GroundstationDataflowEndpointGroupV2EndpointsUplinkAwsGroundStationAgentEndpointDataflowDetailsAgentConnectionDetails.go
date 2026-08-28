// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationdataflowendpointgroupv2


type GroundstationDataflowEndpointGroupV2EndpointsUplinkAwsGroundStationAgentEndpointDataflowDetailsAgentConnectionDetails struct {
	// Socket address of an uplink or downlink agent endpoint with a port range and an optional mtu.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/groundstation_dataflow_endpoint_group_v2#agent_ip_and_port_address GroundstationDataflowEndpointGroupV2#agent_ip_and_port_address}
	AgentIpAndPortAddress *GroundstationDataflowEndpointGroupV2EndpointsUplinkAwsGroundStationAgentEndpointDataflowDetailsAgentConnectionDetailsAgentIpAndPortAddress `field:"optional" json:"agentIpAndPortAddress" yaml:"agentIpAndPortAddress"`
	// Socket address of an uplink or downlink agent endpoint with an optional mtu.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/groundstation_dataflow_endpoint_group_v2#ingress_address_and_port GroundstationDataflowEndpointGroupV2#ingress_address_and_port}
	IngressAddressAndPort *GroundstationDataflowEndpointGroupV2EndpointsUplinkAwsGroundStationAgentEndpointDataflowDetailsAgentConnectionDetailsIngressAddressAndPort `field:"optional" json:"ingressAddressAndPort" yaml:"ingressAddressAndPort"`
}

