// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputConfigurationMerge struct {
	// The time window in milliseconds for merging the two input sources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_router_input#merge_recovery_window_milliseconds MediaconnectRouterInput#merge_recovery_window_milliseconds}
	MergeRecoveryWindowMilliseconds *float64 `field:"optional" json:"mergeRecoveryWindowMilliseconds" yaml:"mergeRecoveryWindowMilliseconds"`
	// The ARN of the network interface to use for this merge router input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_router_input#network_interface_arn MediaconnectRouterInput#network_interface_arn}
	NetworkInterfaceArn *string `field:"optional" json:"networkInterfaceArn" yaml:"networkInterfaceArn"`
	// A list of exactly two protocol configurations for the merge input sources. Both must use the same protocol type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_router_input#protocol_configurations MediaconnectRouterInput#protocol_configurations}
	ProtocolConfigurations interface{} `field:"optional" json:"protocolConfigurations" yaml:"protocolConfigurations"`
}

