// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewisegateway


type IotsitewiseGatewayGatewayCapabilitySummaries struct {
	// The JSON document that defines the gateway capability's configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_gateway#capability_configuration IotsitewiseGateway#capability_configuration}
	CapabilityConfiguration *string `field:"optional" json:"capabilityConfiguration" yaml:"capabilityConfiguration"`
	// The namespace of the capability configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_gateway#capability_namespace IotsitewiseGateway#capability_namespace}
	CapabilityNamespace *string `field:"optional" json:"capabilityNamespace" yaml:"capabilityNamespace"`
}

