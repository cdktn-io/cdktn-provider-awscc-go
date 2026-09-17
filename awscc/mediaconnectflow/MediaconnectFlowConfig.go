// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectFlowConfig struct {
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
	// The name of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow#name MediaconnectFlow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The source of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow#source MediaconnectFlow#source}
	Source *MediaconnectFlowSource `field:"required" json:"source" yaml:"source"`
	// The Availability Zone that you want to create the flow in.
	//
	// These options are limited to the Availability Zones within the current AWS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow#availability_zone MediaconnectFlow#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The encoding configuration to apply to the NDI source content when transcoding it to a transport stream (TS) for downstream distribution.
	//
	// You can choose between several predefined encoding profiles based on common use cases.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow#encoding_config MediaconnectFlow#encoding_config}
	EncodingConfig *MediaconnectFlowEncodingConfig `field:"optional" json:"encodingConfig" yaml:"encodingConfig"`
	// Determines the processing capacity and feature set of the flow.
	//
	// Set this optional parameter to LARGE if you want to enable NDI sources or outputs on the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow#flow_size MediaconnectFlow#flow_size}
	FlowSize *string `field:"optional" json:"flowSize" yaml:"flowSize"`
	// The maintenance settings you want to use for the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow#maintenance MediaconnectFlow#maintenance}
	Maintenance *MediaconnectFlowMaintenance `field:"optional" json:"maintenance" yaml:"maintenance"`
	// The media streams associated with the flow.
	//
	// You can associate any of these media streams with sources and outputs on the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow#media_streams MediaconnectFlow#media_streams}
	MediaStreams interface{} `field:"optional" json:"mediaStreams" yaml:"mediaStreams"`
	// Specifies the configuration settings for NDI sources and outputs. Required when the flow includes NDI sources or outputs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow#ndi_config MediaconnectFlow#ndi_config}
	NdiConfig *MediaconnectFlowNdiConfig `field:"optional" json:"ndiConfig" yaml:"ndiConfig"`
	// The source failover config of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow#source_failover_config MediaconnectFlow#source_failover_config}
	SourceFailoverConfig *MediaconnectFlowSourceFailoverConfig `field:"optional" json:"sourceFailoverConfig" yaml:"sourceFailoverConfig"`
	// The source monitoring config of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow#source_monitoring_config MediaconnectFlow#source_monitoring_config}
	SourceMonitoringConfig *MediaconnectFlowSourceMonitoringConfig `field:"optional" json:"sourceMonitoringConfig" yaml:"sourceMonitoringConfig"`
	// Key-value pairs that can be used to tag this flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow#tags MediaconnectFlow#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The VPC interfaces that you added to this flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow#vpc_interfaces MediaconnectFlow#vpc_interfaces}
	VpcInterfaces interface{} `field:"optional" json:"vpcInterfaces" yaml:"vpcInterfaces"`
}

