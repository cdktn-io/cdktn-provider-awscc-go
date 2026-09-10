// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2vpnconnection


type Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptions struct {
	// Options for sending VPN tunnel logs to CloudWatch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_vpn_connection#cloudwatch_log_options Ec2VpnConnection#cloudwatch_log_options}
	CloudwatchLogOptions *Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptionsCloudwatchLogOptions `field:"optional" json:"cloudwatchLogOptions" yaml:"cloudwatchLogOptions"`
}

