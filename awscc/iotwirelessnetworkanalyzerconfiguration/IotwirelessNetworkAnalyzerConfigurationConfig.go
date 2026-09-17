// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelessnetworkanalyzerconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotwirelessNetworkAnalyzerConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Name of the network analyzer configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotwireless_network_analyzer_configuration#name IotwirelessNetworkAnalyzerConfiguration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotwireless_network_analyzer_configuration#description IotwirelessNetworkAnalyzerConfiguration#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotwireless_network_analyzer_configuration#tags IotwirelessNetworkAnalyzerConfiguration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Trace content for your wireless gateway and wireless device resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotwireless_network_analyzer_configuration#trace_content IotwirelessNetworkAnalyzerConfiguration#trace_content}
	TraceContent *IotwirelessNetworkAnalyzerConfigurationTraceContent `field:"optional" json:"traceContent" yaml:"traceContent"`
	// List of wireless gateway resources that have been added to the network analyzer configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotwireless_network_analyzer_configuration#wireless_devices IotwirelessNetworkAnalyzerConfiguration#wireless_devices}
	WirelessDevices *[]*string `field:"optional" json:"wirelessDevices" yaml:"wirelessDevices"`
	// List of wireless gateway resources that have been added to the network analyzer configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotwireless_network_analyzer_configuration#wireless_gateways IotwirelessNetworkAnalyzerConfiguration#wireless_gateways}
	WirelessGateways *[]*string `field:"optional" json:"wirelessGateways" yaml:"wirelessGateways"`
}

