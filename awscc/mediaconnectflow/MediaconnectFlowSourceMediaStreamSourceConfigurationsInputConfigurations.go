// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow


type MediaconnectFlowSourceMediaStreamSourceConfigurationsInputConfigurations struct {
	// The port that the flow listens on for an incoming media stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow#input_port MediaconnectFlow#input_port}
	InputPort *float64 `field:"optional" json:"inputPort" yaml:"inputPort"`
	// The VPC interface where the media stream comes in from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow#interface MediaconnectFlow#interface}
	Interface *MediaconnectFlowSourceMediaStreamSourceConfigurationsInputConfigurationsInterface `field:"optional" json:"interface" yaml:"interface"`
}

