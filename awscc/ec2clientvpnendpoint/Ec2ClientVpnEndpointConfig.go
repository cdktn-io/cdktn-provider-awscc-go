// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2clientvpnendpoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2ClientVpnEndpointConfig struct {
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
	// Information about the authentication method to be used to authenticate clients.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#authentication_options Ec2ClientVpnEndpoint#authentication_options}
	AuthenticationOptions interface{} `field:"required" json:"authenticationOptions" yaml:"authenticationOptions"`
	// Information about the client connection logging options.
	//
	// If you enable client connection logging, data about client connections is sent to a Cloudwatch Logs log stream. The following information is logged: client connection requests, client connection results (successful and unsuccessful), reasons for unsuccessful client connection requests, client connection termination time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#connection_log_options Ec2ClientVpnEndpoint#connection_log_options}
	ConnectionLogOptions *Ec2ClientVpnEndpointConnectionLogOptions `field:"required" json:"connectionLogOptions" yaml:"connectionLogOptions"`
	// The ARN of the server certificate. For more information, see the AWS Certificate Manager User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#server_certificate_arn Ec2ClientVpnEndpoint#server_certificate_arn}
	ServerCertificateArn *string `field:"required" json:"serverCertificateArn" yaml:"serverCertificateArn"`
	// The Cedar authorization policy to associate with the Client VPN endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#authorization_policy Ec2ClientVpnEndpoint#authorization_policy}
	AuthorizationPolicy *Ec2ClientVpnEndpointAuthorizationPolicy `field:"optional" json:"authorizationPolicy" yaml:"authorizationPolicy"`
	// The IPv4 address range, in CIDR notation, from which to assign client IP addresses.
	//
	// The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. Client CIDR range must have a size of at least /22 and must not be greater than /12.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#client_cidr_block Ec2ClientVpnEndpoint#client_cidr_block}
	ClientCidrBlock *string `field:"optional" json:"clientCidrBlock" yaml:"clientCidrBlock"`
	// The options for managing connection authorization for new client connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#client_connect_options Ec2ClientVpnEndpoint#client_connect_options}
	ClientConnectOptions *Ec2ClientVpnEndpointClientConnectOptions `field:"optional" json:"clientConnectOptions" yaml:"clientConnectOptions"`
	// Options for enabling a customizable text banner that will be displayed on AWS provided clients when a VPN session is established.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#client_login_banner_options Ec2ClientVpnEndpoint#client_login_banner_options}
	ClientLoginBannerOptions *Ec2ClientVpnEndpointClientLoginBannerOptions `field:"optional" json:"clientLoginBannerOptions" yaml:"clientLoginBannerOptions"`
	// Client route enforcement is a feature of the Client VPN service that helps enforce administrator defined routes on devices connected through the VPN.
	//
	// This feature helps improve your security posture by ensuring that network traffic originating from a connected client is not inadvertently sent outside the VPN tunnel.
	// Client route enforcement works by monitoring the route table of a connected device for routing policy changes to the VPN connection. If the feature detects any VPN routing policy modifications, it will automatically force an update to the route table, reverting it back to the expected route configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#client_route_enforcement_options Ec2ClientVpnEndpoint#client_route_enforcement_options}
	ClientRouteEnforcementOptions *Ec2ClientVpnEndpointClientRouteEnforcementOptions `field:"optional" json:"clientRouteEnforcementOptions" yaml:"clientRouteEnforcementOptions"`
	// A brief description of the Client VPN endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#description Ec2ClientVpnEndpoint#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The device-posture trust providers to associate with the Client VPN endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#device_posture_options Ec2ClientVpnEndpoint#device_posture_options}
	DevicePostureOptions *Ec2ClientVpnEndpointDevicePostureOptions `field:"optional" json:"devicePostureOptions" yaml:"devicePostureOptions"`
	// Indicates whether the client VPN session is disconnected after the maximum sessionTimeoutHours is reached.
	//
	// If true, users are prompted to reconnect client VPN. If false, client VPN attempts to reconnect automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#disconnect_on_session_timeout Ec2ClientVpnEndpoint#disconnect_on_session_timeout}
	DisconnectOnSessionTimeout interface{} `field:"optional" json:"disconnectOnSessionTimeout" yaml:"disconnectOnSessionTimeout"`
	// Information about the DNS servers to be used for DNS resolution.
	//
	// A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address configured on the device is used for the DNS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#dns_servers Ec2ClientVpnEndpoint#dns_servers}
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// The IP address type of the Client VPN endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#endpoint_ip_address_type Ec2ClientVpnEndpoint#endpoint_ip_address_type}
	EndpointIpAddressType *string `field:"optional" json:"endpointIpAddressType" yaml:"endpointIpAddressType"`
	// The IDs of one or more security groups to apply to the target network.
	//
	// You must also specify the ID of the VPC that contains the security groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#security_group_ids Ec2ClientVpnEndpoint#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Specify whether to enable the self-service portal for the Client VPN endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#self_service_portal Ec2ClientVpnEndpoint#self_service_portal}
	SelfServicePortal *string `field:"optional" json:"selfServicePortal" yaml:"selfServicePortal"`
	// The maximum VPN session duration time in hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#session_timeout_hours Ec2ClientVpnEndpoint#session_timeout_hours}
	SessionTimeoutHours *float64 `field:"optional" json:"sessionTimeoutHours" yaml:"sessionTimeoutHours"`
	// Indicates whether split-tunnel is enabled on the AWS Client VPN endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#split_tunnel Ec2ClientVpnEndpoint#split_tunnel}
	SplitTunnel interface{} `field:"optional" json:"splitTunnel" yaml:"splitTunnel"`
	// The tags to apply to the Client VPN endpoint during creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#tag_specifications Ec2ClientVpnEndpoint#tag_specifications}
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// The IP address type of the Client VPN endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#traffic_ip_address_type Ec2ClientVpnEndpoint#traffic_ip_address_type}
	TrafficIpAddressType *string `field:"optional" json:"trafficIpAddressType" yaml:"trafficIpAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#transit_gateway_configuration Ec2ClientVpnEndpoint#transit_gateway_configuration}.
	TransitGatewayConfiguration *Ec2ClientVpnEndpointTransitGatewayConfiguration `field:"optional" json:"transitGatewayConfiguration" yaml:"transitGatewayConfiguration"`
	// The transport protocol to be used by the VPN session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#transport_protocol Ec2ClientVpnEndpoint#transport_protocol}
	TransportProtocol *string `field:"optional" json:"transportProtocol" yaml:"transportProtocol"`
	// The ID of the VPC to associate with the Client VPN endpoint.
	//
	// If no security group IDs are specified in the request, the default security group for the VPC is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#vpc_id Ec2ClientVpnEndpoint#vpc_id}
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// The port number to assign to the Client VPN endpoint for TCP and UDP traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#vpn_port Ec2ClientVpnEndpoint#vpn_port}
	VpnPort *float64 `field:"optional" json:"vpnPort" yaml:"vpnPort"`
}

