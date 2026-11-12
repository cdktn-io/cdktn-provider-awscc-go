// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmonitorprobe

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkmonitorProbeConfig struct {
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
	// The literal IPv4 or IPv6 address the probe sends traffic to.
	//
	// The service rejects a domain name and a CIDR block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmonitor_probe#destination NetworkmonitorProbe#destination}
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// The name of the monitor the probe belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmonitor_probe#monitor_name NetworkmonitorProbe#monitor_name}
	MonitorName *string `field:"required" json:"monitorName" yaml:"monitorName"`
	// The protocol the probe uses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmonitor_probe#protocol NetworkmonitorProbe#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The ARN of the subnet the probe sends traffic from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmonitor_probe#source_arn NetworkmonitorProbe#source_arn}
	SourceArn *string `field:"required" json:"sourceArn" yaml:"sourceArn"`
	// The port the probe sends traffic to. Applies to the TCP protocol only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmonitor_probe#destination_port NetworkmonitorProbe#destination_port}
	DestinationPort *float64 `field:"optional" json:"destinationPort" yaml:"destinationPort"`
	// The size, in bytes, of the packets the probe sends.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmonitor_probe#packet_size NetworkmonitorProbe#packet_size}
	PacketSize *float64 `field:"optional" json:"packetSize" yaml:"packetSize"`
	// The tags for the probe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/networkmonitor_probe#tags NetworkmonitorProbe#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

