// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputConfigurationFailover struct {
	// The ARN of the network interface to use for this failover router input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_input#network_interface_arn MediaconnectRouterInput#network_interface_arn}
	NetworkInterfaceArn *string `field:"optional" json:"networkInterfaceArn" yaml:"networkInterfaceArn"`
	// The index (0 or 1) that specifies which source in the protocol configurations list is currently active.
	//
	// Used to control which of the two failover sources is currently selected. This field is ignored when sourcePriorityMode is set to NO_PRIORITY
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_input#primary_source_index MediaconnectRouterInput#primary_source_index}
	PrimarySourceIndex *float64 `field:"optional" json:"primarySourceIndex" yaml:"primarySourceIndex"`
	// A list of exactly two protocol configurations for the failover input sources. Both must use the same protocol type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_input#protocol_configurations MediaconnectRouterInput#protocol_configurations}
	ProtocolConfigurations interface{} `field:"optional" json:"protocolConfigurations" yaml:"protocolConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_input#source_priority_mode MediaconnectRouterInput#source_priority_mode}.
	SourcePriorityMode *string `field:"optional" json:"sourcePriorityMode" yaml:"sourcePriorityMode"`
}

