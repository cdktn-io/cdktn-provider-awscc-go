// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdanetworkconnector


type LambdaNetworkConnectorConfiguration struct {
	// The VPC egress configuration for the network connector.
	//
	// Specifies the subnets, security groups, and network protocol for routing outbound traffic through your VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_network_connector#vpc_egress_configuration LambdaNetworkConnector#vpc_egress_configuration}
	VpcEgressConfiguration *LambdaNetworkConnectorConfigurationVpcEgressConfiguration `field:"required" json:"vpcEgressConfiguration" yaml:"vpcEgressConfiguration"`
}

