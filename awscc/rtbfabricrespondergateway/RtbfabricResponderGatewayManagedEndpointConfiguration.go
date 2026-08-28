// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabricrespondergateway


type RtbfabricResponderGatewayManagedEndpointConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_responder_gateway#auto_scaling_groups_configuration RtbfabricResponderGateway#auto_scaling_groups_configuration}.
	AutoScalingGroupsConfiguration *RtbfabricResponderGatewayManagedEndpointConfigurationAutoScalingGroupsConfiguration `field:"optional" json:"autoScalingGroupsConfiguration" yaml:"autoScalingGroupsConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_responder_gateway#eks_endpoints_configuration RtbfabricResponderGateway#eks_endpoints_configuration}.
	EksEndpointsConfiguration *RtbfabricResponderGatewayManagedEndpointConfigurationEksEndpointsConfiguration `field:"optional" json:"eksEndpointsConfiguration" yaml:"eksEndpointsConfiguration"`
}

