// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationdataflowendpointgroupv2


type GroundstationDataflowEndpointGroupV2EndpointsUplinkAwsGroundStationAgentEndpoint struct {
	// Dataflow details for uplink.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/groundstation_dataflow_endpoint_group_v2#dataflow_details GroundstationDataflowEndpointGroupV2#dataflow_details}
	DataflowDetails *GroundstationDataflowEndpointGroupV2EndpointsUplinkAwsGroundStationAgentEndpointDataflowDetails `field:"optional" json:"dataflowDetails" yaml:"dataflowDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/groundstation_dataflow_endpoint_group_v2#name GroundstationDataflowEndpointGroupV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

