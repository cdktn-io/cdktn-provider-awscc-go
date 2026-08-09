// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflowsource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectFlowSourceAConfig struct {
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
	// A description for the source.
	//
	// This value is not used or seen outside of the current AWS Elemental MediaConnect account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_source#description MediaconnectFlowSourceA#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The ARN of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_source#flow_arn MediaconnectFlowSourceA#flow_arn}
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// The name of the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_source#name MediaconnectFlowSourceA#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of encryption that is used on the content ingested from this source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_source#decryption MediaconnectFlowSourceA#decryption}
	Decryption *MediaconnectFlowSourceDecryptionA `field:"optional" json:"decryption" yaml:"decryption"`
	// The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account.
	//
	// The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_source#entitlement_arn MediaconnectFlowSourceA#entitlement_arn}
	EntitlementArn *string `field:"optional" json:"entitlementArn" yaml:"entitlementArn"`
	// The source configuration for cloud flows receiving a stream from a bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_source#gateway_bridge_source MediaconnectFlowSourceA#gateway_bridge_source}
	GatewayBridgeSource *MediaconnectFlowSourceGatewayBridgeSourceA `field:"optional" json:"gatewayBridgeSource" yaml:"gatewayBridgeSource"`
	// The port that the flow will be listening on for incoming content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_source#ingest_port MediaconnectFlowSourceA#ingest_port}
	IngestPort *float64 `field:"optional" json:"ingestPort" yaml:"ingestPort"`
	// The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_source#max_bitrate MediaconnectFlowSourceA#max_bitrate}
	MaxBitrate *float64 `field:"optional" json:"maxBitrate" yaml:"maxBitrate"`
	// The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_source#max_latency MediaconnectFlowSourceA#max_latency}
	MaxLatency *float64 `field:"optional" json:"maxLatency" yaml:"maxLatency"`
	// The minimum latency in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_source#min_latency MediaconnectFlowSourceA#min_latency}
	MinLatency *float64 `field:"optional" json:"minLatency" yaml:"minLatency"`
	// The protocol that is used by the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_source#protocol MediaconnectFlowSourceA#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The port that the flow uses to send outbound requests to initiate connection with the sender for fujitsu-qos protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_source#sender_control_port MediaconnectFlowSourceA#sender_control_port}
	SenderControlPort *float64 `field:"optional" json:"senderControlPort" yaml:"senderControlPort"`
	// The IP address that the flow communicates with to initiate connection with the sender for fujitsu-qos protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_source#sender_ip_address MediaconnectFlowSourceA#sender_ip_address}
	SenderIpAddress *string `field:"optional" json:"senderIpAddress" yaml:"senderIpAddress"`
	// Source IP or domain name for SRT-caller protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_source#source_listener_address MediaconnectFlowSourceA#source_listener_address}
	SourceListenerAddress *string `field:"optional" json:"sourceListenerAddress" yaml:"sourceListenerAddress"`
	// Source port for SRT-caller protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_source#source_listener_port MediaconnectFlowSourceA#source_listener_port}
	SourceListenerPort *float64 `field:"optional" json:"sourceListenerPort" yaml:"sourceListenerPort"`
	// The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_source#stream_id MediaconnectFlowSourceA#stream_id}
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
	// Key-value pairs that can be used to tag and organize this flow source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_source#tags MediaconnectFlowSourceA#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The name of the VPC Interface this Source is configured with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_source#vpc_interface_name MediaconnectFlowSourceA#vpc_interface_name}
	VpcInterfaceName *string `field:"optional" json:"vpcInterfaceName" yaml:"vpcInterfaceName"`
	// The range of IP addresses that should be allowed to contribute content to your source.
	//
	// These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_flow_source#whitelist_cidr MediaconnectFlowSourceA#whitelist_cidr}
	WhitelistCidr *string `field:"optional" json:"whitelistCidr" yaml:"whitelistCidr"`
}

