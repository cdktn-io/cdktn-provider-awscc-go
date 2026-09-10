// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput


type MediaconnectRouterOutputConfigurationMediaConnectFlow struct {
	// The configuration that defines how content is encrypted during transit between the MediaConnect router and a MediaConnect flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_output#destination_transit_encryption MediaconnectRouterOutput#destination_transit_encryption}
	DestinationTransitEncryption *MediaconnectRouterOutputConfigurationMediaConnectFlowDestinationTransitEncryption `field:"optional" json:"destinationTransitEncryption" yaml:"destinationTransitEncryption"`
	// The ARN of the flow to connect to this router output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_output#flow_arn MediaconnectRouterOutput#flow_arn}
	FlowArn *string `field:"optional" json:"flowArn" yaml:"flowArn"`
	// The ARN of the flow source to connect to this router output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_output#flow_source_arn MediaconnectRouterOutput#flow_source_arn}
	FlowSourceArn *string `field:"optional" json:"flowSourceArn" yaml:"flowSourceArn"`
}

