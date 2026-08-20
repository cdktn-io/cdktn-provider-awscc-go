// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationdataflowendpointgroupv2


type GroundstationDataflowEndpointGroupV2Endpoints struct {
	// Information about DownlinkAwsGroundStationAgentEndpoint used for create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/groundstation_dataflow_endpoint_group_v2#downlink_aws_ground_station_agent_endpoint GroundstationDataflowEndpointGroupV2#downlink_aws_ground_station_agent_endpoint}
	DownlinkAwsGroundStationAgentEndpoint *GroundstationDataflowEndpointGroupV2EndpointsDownlinkAwsGroundStationAgentEndpoint `field:"optional" json:"downlinkAwsGroundStationAgentEndpoint" yaml:"downlinkAwsGroundStationAgentEndpoint"`
	// Information about UplinkAwsGroundStationAgentEndpoint used for create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/groundstation_dataflow_endpoint_group_v2#uplink_aws_ground_station_agent_endpoint GroundstationDataflowEndpointGroupV2#uplink_aws_ground_station_agent_endpoint}
	UplinkAwsGroundStationAgentEndpoint *GroundstationDataflowEndpointGroupV2EndpointsUplinkAwsGroundStationAgentEndpoint `field:"optional" json:"uplinkAwsGroundStationAgentEndpoint" yaml:"uplinkAwsGroundStationAgentEndpoint"`
}

