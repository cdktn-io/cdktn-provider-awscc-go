// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsgateway


type EcsGatewayConfigurationHostSharedAlb struct {
	// The access log configuration for the gateway ALB. Access logging is enabled by the presence of a destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#access_logs EcsGateway#access_logs}
	AccessLogs *EcsGatewayConfigurationHostSharedAlbAccessLogs `field:"optional" json:"accessLogs" yaml:"accessLogs"`
	// The HTTP desync mitigation mode for the ALB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#http_desync EcsGateway#http_desync}
	HttpDesync *string `field:"optional" json:"httpDesync" yaml:"httpDesync"`
	// The maximum number of ALBs to provision for the gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#max_alb_count EcsGateway#max_alb_count}
	MaxAlbCount *float64 `field:"optional" json:"maxAlbCount" yaml:"maxAlbCount"`
	// The network configuration for the gateway ALB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#network_configuration EcsGateway#network_configuration}
	NetworkConfiguration *EcsGatewayConfigurationHostSharedAlbNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The port on which the ALB listens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#port EcsGateway#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol for the ALB listener.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#protocol EcsGateway#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The scheme for the ALB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#scheme EcsGateway#scheme}
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// The test traffic rules for the gateway ALB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#test_traffic_rules EcsGateway#test_traffic_rules}
	TestTrafficRules *EcsGatewayConfigurationHostSharedAlbTestTrafficRules `field:"optional" json:"testTrafficRules" yaml:"testTrafficRules"`
	// The ARN of the WAFv2 web ACL to associate with the ALB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#web_acl_arn EcsGateway#web_acl_arn}
	WebAclArn *string `field:"optional" json:"webAclArn" yaml:"webAclArn"`
}

