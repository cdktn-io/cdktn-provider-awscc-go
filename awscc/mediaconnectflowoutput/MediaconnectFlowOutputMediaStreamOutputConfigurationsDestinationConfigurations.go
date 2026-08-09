// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflowoutput


type MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurations struct {
	// The IP address where contents of the media stream will be sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_output#destination_ip MediaconnectFlowOutput#destination_ip}
	DestinationIp *string `field:"optional" json:"destinationIp" yaml:"destinationIp"`
	// The port to use when the content of the media stream is distributed to the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_output#destination_port MediaconnectFlowOutput#destination_port}
	DestinationPort *float64 `field:"optional" json:"destinationPort" yaml:"destinationPort"`
	// The VPC interface that is used for the media stream associated with the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_output#interface MediaconnectFlowOutput#interface}
	Interface *MediaconnectFlowOutputMediaStreamOutputConfigurationsDestinationConfigurationsInterface `field:"optional" json:"interface" yaml:"interface"`
}

