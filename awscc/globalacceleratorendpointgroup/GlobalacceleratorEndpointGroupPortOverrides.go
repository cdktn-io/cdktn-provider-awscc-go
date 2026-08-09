// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package globalacceleratorendpointgroup


type GlobalacceleratorEndpointGroupPortOverrides struct {
	// A network port number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/globalaccelerator_endpoint_group#endpoint_port GlobalacceleratorEndpointGroup#endpoint_port}
	EndpointPort *float64 `field:"optional" json:"endpointPort" yaml:"endpointPort"`
	// A network port number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/globalaccelerator_endpoint_group#listener_port GlobalacceleratorEndpointGroup#listener_port}
	ListenerPort *float64 `field:"optional" json:"listenerPort" yaml:"listenerPort"`
}

