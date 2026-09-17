// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationdataflowendpointgroupv2


type GroundstationDataflowEndpointGroupV2EndpointsDownlinkAwsGroundStationAgentEndpointDataflowDetailsAgentConnectionDetailsAgentIpAndPortAddressSocketAddress struct {
	// IPv4 socket address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_dataflow_endpoint_group_v2#name GroundstationDataflowEndpointGroupV2#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Port range of a socket address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_dataflow_endpoint_group_v2#port_range GroundstationDataflowEndpointGroupV2#port_range}
	PortRange *GroundstationDataflowEndpointGroupV2EndpointsDownlinkAwsGroundStationAgentEndpointDataflowDetailsAgentConnectionDetailsAgentIpAndPortAddressSocketAddressPortRange `field:"optional" json:"portRange" yaml:"portRange"`
}

