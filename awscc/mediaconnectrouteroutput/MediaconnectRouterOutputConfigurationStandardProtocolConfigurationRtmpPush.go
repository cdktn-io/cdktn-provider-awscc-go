// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput


type MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPush struct {
	// The name of the RTMP application on the destination server.
	//
	// Together with the stream name, the application name forms the RTMP URL path, in the pattern rtmp://destinationAddress/applicationName/streamName.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_router_output#application_name MediaconnectRouterOutput#application_name}
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// The IP address or hostname of the destination RTMP server that the router output pushes the stream to.
	//
	// Provide only the server address; specify the application and stream names separately.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_router_output#destination_address MediaconnectRouterOutput#destination_address}
	DestinationAddress *string `field:"optional" json:"destinationAddress" yaml:"destinationAddress"`
	// The TCP port on the destination RTMP server.
	//
	// For RTMP, valid values range from 1024 to 65535. For RTMPS (RTMP over TLS), valid values are 443 or 1024 to 65535. RTMP typically uses port 1935, and RTMPS typically uses port 443.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_router_output#destination_port MediaconnectRouterOutput#destination_port}
	DestinationPort *float64 `field:"optional" json:"destinationPort" yaml:"destinationPort"`
	// The name of the RTMP stream that the output publishes to the destination application.
	//
	// The stream name forms the final segment of the RTMP URL path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_router_output#stream_name MediaconnectRouterOutput#stream_name}
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
	// The Transport Layer Security (TLS) encryption settings used to establish a secure connection to a destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_router_output#tls_encryption MediaconnectRouterOutput#tls_encryption}
	TlsEncryption *MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushTlsEncryption `field:"optional" json:"tlsEncryption" yaml:"tlsEncryption"`
}

