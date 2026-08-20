// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow


type MediaconnectFlowSource struct {
	// The type of decryption that is used on the content ingested from this source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#decryption MediaconnectFlow#decryption}
	Decryption *MediaconnectFlowSourceDecryption `field:"optional" json:"decryption" yaml:"decryption"`
	// A description for the source.
	//
	// This value is not used or seen outside of the current AWS Elemental MediaConnect account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#description MediaconnectFlow#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account.
	//
	// The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#entitlement_arn MediaconnectFlow#entitlement_arn}
	EntitlementArn *string `field:"optional" json:"entitlementArn" yaml:"entitlementArn"`
	// The source configuration for cloud flows receiving a stream from a bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#gateway_bridge_source MediaconnectFlow#gateway_bridge_source}
	GatewayBridgeSource *MediaconnectFlowSourceGatewayBridgeSource `field:"optional" json:"gatewayBridgeSource" yaml:"gatewayBridgeSource"`
	// The port that the flow will be listening on for incoming content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#ingest_port MediaconnectFlow#ingest_port}
	IngestPort *float64 `field:"optional" json:"ingestPort" yaml:"ingestPort"`
	// The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#max_bitrate MediaconnectFlow#max_bitrate}
	MaxBitrate *float64 `field:"optional" json:"maxBitrate" yaml:"maxBitrate"`
	// The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#max_latency MediaconnectFlow#max_latency}
	MaxLatency *float64 `field:"optional" json:"maxLatency" yaml:"maxLatency"`
	// The size of the buffer (in milliseconds) to use to sync incoming source data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#max_sync_buffer MediaconnectFlow#max_sync_buffer}
	MaxSyncBuffer *float64 `field:"optional" json:"maxSyncBuffer" yaml:"maxSyncBuffer"`
	// The media stream that is associated with the source, and the parameters for that association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#media_stream_source_configurations MediaconnectFlow#media_stream_source_configurations}
	MediaStreamSourceConfigurations interface{} `field:"optional" json:"mediaStreamSourceConfigurations" yaml:"mediaStreamSourceConfigurations"`
	// The minimum latency in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#min_latency MediaconnectFlow#min_latency}
	MinLatency *float64 `field:"optional" json:"minLatency" yaml:"minLatency"`
	// The name of the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#name MediaconnectFlow#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The protocol that is used by the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#protocol MediaconnectFlow#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The port that the flow uses to send outbound requests to initiate connection with the sender for fujitsu-qos protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#sender_control_port MediaconnectFlow#sender_control_port}
	SenderControlPort *float64 `field:"optional" json:"senderControlPort" yaml:"senderControlPort"`
	// The IP address that the flow communicates with to initiate connection with the sender for fujitsu-qos protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#sender_ip_address MediaconnectFlow#sender_ip_address}
	SenderIpAddress *string `field:"optional" json:"senderIpAddress" yaml:"senderIpAddress"`
	// Source IP or domain name for SRT-caller protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#source_listener_address MediaconnectFlow#source_listener_address}
	SourceListenerAddress *string `field:"optional" json:"sourceListenerAddress" yaml:"sourceListenerAddress"`
	// Source port for SRT-caller protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#source_listener_port MediaconnectFlow#source_listener_port}
	SourceListenerPort *float64 `field:"optional" json:"sourceListenerPort" yaml:"sourceListenerPort"`
	// The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#stream_id MediaconnectFlow#stream_id}
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
	// The name of the VPC Interface this Source is configured with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#vpc_interface_name MediaconnectFlow#vpc_interface_name}
	VpcInterfaceName *string `field:"optional" json:"vpcInterfaceName" yaml:"vpcInterfaceName"`
	// The range of IP addresses that should be allowed to contribute content to your source.
	//
	// These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#whitelist_cidr MediaconnectFlow#whitelist_cidr}
	WhitelistCidr *string `field:"optional" json:"whitelistCidr" yaml:"whitelistCidr"`
}

