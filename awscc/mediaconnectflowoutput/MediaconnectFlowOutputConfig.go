// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflowoutput

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectFlowOutputConfig struct {
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
	// The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#flow_arn MediaconnectFlowOutput#flow_arn}
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// The range of IP addresses that should be allowed to initiate output requests to this flow.
	//
	// These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#cidr_allow_list MediaconnectFlowOutput#cidr_allow_list}
	CidrAllowList *[]*string `field:"optional" json:"cidrAllowList" yaml:"cidrAllowList"`
	// A description of the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#description MediaconnectFlowOutput#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The address where you want to send the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#destination MediaconnectFlowOutput#destination}
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// The type of key used for the encryption.
	//
	// If no keyType is provided, the service will use the default setting (static-key).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#encryption MediaconnectFlowOutput#encryption}
	Encryption *MediaconnectFlowOutputEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#max_latency MediaconnectFlowOutput#max_latency}
	MaxLatency *float64 `field:"optional" json:"maxLatency" yaml:"maxLatency"`
	// The definition for each media stream that is associated with the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#media_stream_output_configurations MediaconnectFlowOutput#media_stream_output_configurations}
	MediaStreamOutputConfigurations interface{} `field:"optional" json:"mediaStreamOutputConfigurations" yaml:"mediaStreamOutputConfigurations"`
	// The minimum latency in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#min_latency MediaconnectFlowOutput#min_latency}
	MinLatency *float64 `field:"optional" json:"minLatency" yaml:"minLatency"`
	// The name of the output. This value must be unique within the current flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#name MediaconnectFlowOutput#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The timecode source for the NDI output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#ndi_output_timecode_source MediaconnectFlowOutput#ndi_output_timecode_source}
	NdiOutputTimecodeSource *string `field:"optional" json:"ndiOutputTimecodeSource" yaml:"ndiOutputTimecodeSource"`
	// A suffix for the names of the NDI sources that the flow creates.
	//
	// If a custom name isn't specified, MediaConnect uses the output name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#ndi_program_name MediaconnectFlowOutput#ndi_program_name}
	NdiProgramName *string `field:"optional" json:"ndiProgramName" yaml:"ndiProgramName"`
	// A quality setting for the NDI Speed HQ encoder.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#ndi_speed_hq_quality MediaconnectFlowOutput#ndi_speed_hq_quality}
	NdiSpeedHqQuality *float64 `field:"optional" json:"ndiSpeedHqQuality" yaml:"ndiSpeedHqQuality"`
	// An indication of whether the output should transmit data or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#output_status MediaconnectFlowOutput#output_status}
	OutputStatus *string `field:"optional" json:"outputStatus" yaml:"outputStatus"`
	// The port to use when content is distributed to this output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#port MediaconnectFlowOutput#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol that is used by the source or output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#protocol MediaconnectFlowOutput#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The remote ID for the Zixi-pull stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#remote_id MediaconnectFlowOutput#remote_id}
	RemoteId *string `field:"optional" json:"remoteId" yaml:"remoteId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#router_integration_state MediaconnectFlowOutput#router_integration_state}.
	RouterIntegrationState *string `field:"optional" json:"routerIntegrationState" yaml:"routerIntegrationState"`
	// The configuration that defines how content is encrypted during transit between the MediaConnect router and a MediaConnect flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#router_integration_transit_encryption MediaconnectFlowOutput#router_integration_transit_encryption}
	RouterIntegrationTransitEncryption *MediaconnectFlowOutputRouterIntegrationTransitEncryption `field:"optional" json:"routerIntegrationTransitEncryption" yaml:"routerIntegrationTransitEncryption"`
	// The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#smoothing_latency MediaconnectFlowOutput#smoothing_latency}
	SmoothingLatency *float64 `field:"optional" json:"smoothingLatency" yaml:"smoothingLatency"`
	// The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#stream_id MediaconnectFlowOutput#stream_id}
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
	// Key-value pairs that can be used to tag and organize this flow output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#tags MediaconnectFlowOutput#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The name of the VPC interface attachment to use for this output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow_output#vpc_interface_attachment MediaconnectFlowOutput#vpc_interface_attachment}
	VpcInterfaceAttachment *MediaconnectFlowOutputVpcInterfaceAttachment `field:"optional" json:"vpcInterfaceAttachment" yaml:"vpcInterfaceAttachment"`
}

