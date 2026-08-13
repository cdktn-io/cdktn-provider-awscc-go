// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2vpnconnection


type Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptions struct {
	// Specifies whether to enable BGP logging for the VPN connection. Default value is ``False``.  Valid values: ``True`` | ``False``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#bgp_log_enabled Ec2VpnConnection#bgp_log_enabled}
	BgpLogEnabled interface{} `field:"optional" json:"bgpLogEnabled" yaml:"bgpLogEnabled"`
	// The Amazon Resource Name (ARN) of the CloudWatch log group where BGP logs will be sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#bgp_log_group_arn Ec2VpnConnection#bgp_log_group_arn}
	BgpLogGroupArn *string `field:"optional" json:"bgpLogGroupArn" yaml:"bgpLogGroupArn"`
	// The desired output format for BGP logs to be sent to CloudWatch.
	//
	// Default format is ``json``.
	//  Valid values: ``json`` | ``text``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#bgp_log_output_format Ec2VpnConnection#bgp_log_output_format}
	BgpLogOutputFormat *string `field:"optional" json:"bgpLogOutputFormat" yaml:"bgpLogOutputFormat"`
	// Enable or disable VPN tunnel logging feature. Default value is ``False``.  Valid values: ``True`` | ``False``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#log_enabled Ec2VpnConnection#log_enabled}
	LogEnabled interface{} `field:"optional" json:"logEnabled" yaml:"logEnabled"`
	// The Amazon Resource Name (ARN) of the CloudWatch log group to send logs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#log_group_arn Ec2VpnConnection#log_group_arn}
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
	// Set log format. Default format is ``json``.  Valid values: ``json`` | ``text``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#log_output_format Ec2VpnConnection#log_output_format}
	LogOutputFormat *string `field:"optional" json:"logOutputFormat" yaml:"logOutputFormat"`
}

