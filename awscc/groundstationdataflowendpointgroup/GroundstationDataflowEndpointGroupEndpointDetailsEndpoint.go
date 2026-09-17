// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationdataflowendpointgroup


type GroundstationDataflowEndpointGroupEndpointDetailsEndpoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_dataflow_endpoint_group#address GroundstationDataflowEndpointGroup#address}.
	Address *GroundstationDataflowEndpointGroupEndpointDetailsEndpointAddress `field:"optional" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_dataflow_endpoint_group#mtu GroundstationDataflowEndpointGroup#mtu}.
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_dataflow_endpoint_group#name GroundstationDataflowEndpointGroup#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

