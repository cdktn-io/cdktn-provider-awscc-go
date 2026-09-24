// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationdataflowendpointgroupv2


type GroundstationDataflowEndpointGroupV2EndpointsDownlinkAwsGroundStationAgentEndpointDataflowDetails struct {
	// Connection details for downlink, from ground station to agent, and customer to agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/groundstation_dataflow_endpoint_group_v2#agent_connection_details GroundstationDataflowEndpointGroupV2#agent_connection_details}
	AgentConnectionDetails *GroundstationDataflowEndpointGroupV2EndpointsDownlinkAwsGroundStationAgentEndpointDataflowDetailsAgentConnectionDetails `field:"optional" json:"agentConnectionDetails" yaml:"agentConnectionDetails"`
}

