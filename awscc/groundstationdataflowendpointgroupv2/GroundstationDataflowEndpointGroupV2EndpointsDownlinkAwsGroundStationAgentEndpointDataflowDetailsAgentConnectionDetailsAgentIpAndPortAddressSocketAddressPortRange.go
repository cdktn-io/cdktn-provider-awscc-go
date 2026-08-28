// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationdataflowendpointgroupv2


type GroundstationDataflowEndpointGroupV2EndpointsDownlinkAwsGroundStationAgentEndpointDataflowDetailsAgentConnectionDetailsAgentIpAndPortAddressSocketAddressPortRange struct {
	// A maximum value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/groundstation_dataflow_endpoint_group_v2#maximum GroundstationDataflowEndpointGroupV2#maximum}
	Maximum *float64 `field:"optional" json:"maximum" yaml:"maximum"`
	// A minimum value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/groundstation_dataflow_endpoint_group_v2#minimum GroundstationDataflowEndpointGroupV2#minimum}
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
}

