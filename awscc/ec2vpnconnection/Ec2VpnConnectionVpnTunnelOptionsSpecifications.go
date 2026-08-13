// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2vpnconnection


type Ec2VpnConnectionVpnTunnelOptionsSpecifications struct {
	// The action to take after DPD timeout occurs.
	//
	// Specify ``restart`` to restart the IKE initiation. Specify ``clear`` to end the IKE session.
	//  Valid Values: ``clear`` | ``none`` | ``restart``
	//  Default: ``clear``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#dpd_timeout_action Ec2VpnConnection#dpd_timeout_action}
	DpdTimeoutAction *string `field:"optional" json:"dpdTimeoutAction" yaml:"dpdTimeoutAction"`
	// The number of seconds after which a DPD timeout occurs.
	//
	// Constraints: A value greater than or equal to 30.
	//  Default: ``30``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#dpd_timeout_seconds Ec2VpnConnection#dpd_timeout_seconds}
	DpdTimeoutSeconds *float64 `field:"optional" json:"dpdTimeoutSeconds" yaml:"dpdTimeoutSeconds"`
	// Turn on or off tunnel endpoint lifecycle control feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#enable_tunnel_lifecycle_control Ec2VpnConnection#enable_tunnel_lifecycle_control}
	EnableTunnelLifecycleControl interface{} `field:"optional" json:"enableTunnelLifecycleControl" yaml:"enableTunnelLifecycleControl"`
	// The IKE versions that are permitted for the VPN tunnel.  Valid values: ``ikev1`` | ``ikev2``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#ike_versions Ec2VpnConnection#ike_versions}
	IkeVersions interface{} `field:"optional" json:"ikeVersions" yaml:"ikeVersions"`
	// Options for logging VPN tunnel activity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#log_options Ec2VpnConnection#log_options}
	LogOptions *Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptions `field:"optional" json:"logOptions" yaml:"logOptions"`
	// One or more Diffie-Hellman group numbers that are permitted for the VPN tunnel for phase 1 IKE negotiations.
	//
	// Valid values: ``2`` | ``14`` | ``15`` | ``16`` | ``17`` | ``18`` | ``19`` | ``20`` | ``21`` | ``22`` | ``23`` | ``24``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#phase_1_dh_group_numbers Ec2VpnConnection#phase_1_dh_group_numbers}
	Phase1DhGroupNumbers interface{} `field:"optional" json:"phase1DhGroupNumbers" yaml:"phase1DhGroupNumbers"`
	// One or more encryption algorithms that are permitted for the VPN tunnel for phase 1 IKE negotiations.
	//
	// Valid values: ``AES128`` | ``AES256`` | ``AES128-GCM-16`` | ``AES256-GCM-16``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#phase_1_encryption_algorithms Ec2VpnConnection#phase_1_encryption_algorithms}
	Phase1EncryptionAlgorithms interface{} `field:"optional" json:"phase1EncryptionAlgorithms" yaml:"phase1EncryptionAlgorithms"`
	// One or more integrity algorithms that are permitted for the VPN tunnel for phase 1 IKE negotiations.
	//
	// Valid values: ``SHA1`` | ``SHA2-256`` | ``SHA2-384`` | ``SHA2-512``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#phase_1_integrity_algorithms Ec2VpnConnection#phase_1_integrity_algorithms}
	Phase1IntegrityAlgorithms interface{} `field:"optional" json:"phase1IntegrityAlgorithms" yaml:"phase1IntegrityAlgorithms"`
	// The lifetime for phase 1 of the IKE negotiation, in seconds.
	//
	// Constraints: A value between 900 and 28,800.
	//  Default: ``28800``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#phase_1_lifetime_seconds Ec2VpnConnection#phase_1_lifetime_seconds}
	Phase1LifetimeSeconds *float64 `field:"optional" json:"phase1LifetimeSeconds" yaml:"phase1LifetimeSeconds"`
	// One or more Diffie-Hellman group numbers that are permitted for the VPN tunnel for phase 2 IKE negotiations.
	//
	// Valid values: ``2`` | ``5`` | ``14`` | ``15`` | ``16`` | ``17`` | ``18`` | ``19`` | ``20`` | ``21`` | ``22`` | ``23`` | ``24``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#phase_2_dh_group_numbers Ec2VpnConnection#phase_2_dh_group_numbers}
	Phase2DhGroupNumbers interface{} `field:"optional" json:"phase2DhGroupNumbers" yaml:"phase2DhGroupNumbers"`
	// One or more encryption algorithms that are permitted for the VPN tunnel for phase 2 IKE negotiations.
	//
	// Valid values: ``AES128`` | ``AES256`` | ``AES128-GCM-16`` | ``AES256-GCM-16``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#phase_2_encryption_algorithms Ec2VpnConnection#phase_2_encryption_algorithms}
	Phase2EncryptionAlgorithms interface{} `field:"optional" json:"phase2EncryptionAlgorithms" yaml:"phase2EncryptionAlgorithms"`
	// One or more integrity algorithms that are permitted for the VPN tunnel for phase 2 IKE negotiations.
	//
	// Valid values: ``SHA1`` | ``SHA2-256`` | ``SHA2-384`` | ``SHA2-512``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#phase_2_integrity_algorithms Ec2VpnConnection#phase_2_integrity_algorithms}
	Phase2IntegrityAlgorithms interface{} `field:"optional" json:"phase2IntegrityAlgorithms" yaml:"phase2IntegrityAlgorithms"`
	// The lifetime for phase 2 of the IKE negotiation, in seconds.
	//
	// Constraints: A value between 900 and 3,600. The value must be less than the value for ``Phase1LifetimeSeconds``.
	//  Default: ``3600``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#phase_2_lifetime_seconds Ec2VpnConnection#phase_2_lifetime_seconds}
	Phase2LifetimeSeconds *float64 `field:"optional" json:"phase2LifetimeSeconds" yaml:"phase2LifetimeSeconds"`
	// The pre-shared key (PSK) to establish initial authentication between the virtual private gateway and customer gateway.
	//
	// Constraints: Allowed characters are alphanumeric characters, periods (.), and underscores (_). Must be between 8 and 64 characters in length and cannot start with zero (0).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#pre_shared_key Ec2VpnConnection#pre_shared_key}
	PreSharedKey *string `field:"optional" json:"preSharedKey" yaml:"preSharedKey"`
	// The percentage of the rekey window (determined by ``RekeyMarginTimeSeconds``) during which the rekey time is randomly selected.
	//
	// Constraints: A value between 0 and 100.
	//  Default: ``100``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#rekey_fuzz_percentage Ec2VpnConnection#rekey_fuzz_percentage}
	RekeyFuzzPercentage *float64 `field:"optional" json:"rekeyFuzzPercentage" yaml:"rekeyFuzzPercentage"`
	// The margin time, in seconds, before the phase 2 lifetime expires, during which the AWS side of the VPN connection performs an IKE rekey.
	//
	// The exact time of the rekey is randomly selected based on the value for ``RekeyFuzzPercentage``.
	//  Constraints: A value between 60 and half of ``Phase2LifetimeSeconds``.
	//  Default: ``270``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#rekey_margin_time_seconds Ec2VpnConnection#rekey_margin_time_seconds}
	RekeyMarginTimeSeconds *float64 `field:"optional" json:"rekeyMarginTimeSeconds" yaml:"rekeyMarginTimeSeconds"`
	// The number of packets in an IKE replay window.  Constraints: A value between 64 and 2048.  Default: ``1024``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#replay_window_size Ec2VpnConnection#replay_window_size}
	ReplayWindowSize *float64 `field:"optional" json:"replayWindowSize" yaml:"replayWindowSize"`
	// The action to take when the establishing the tunnel for the VPN connection.
	//
	// By default, your customer gateway device must initiate the IKE negotiation and bring up the tunnel. Specify ``start`` for AWS to initiate the IKE negotiation.
	//  Valid Values: ``add`` | ``start``
	//  Default: ``add``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#startup_action Ec2VpnConnection#startup_action}
	StartupAction *string `field:"optional" json:"startupAction" yaml:"startupAction"`
	// The range of inside IP addresses for the tunnel.
	//
	// Any specified CIDR blocks must be unique across all VPN connections that use the same virtual private gateway.
	//  Constraints: A size /30 CIDR block from the ``169.254.0.0/16`` range. The following CIDR blocks are reserved and cannot be used:
	//   +   ``169.254.0.0/30``
	//   +   ``169.254.1.0/30``
	//   +   ``169.254.2.0/30``
	//   +   ``169.254.3.0/30``
	//   +   ``169.254.4.0/30``
	//   +   ``169.254.5.0/30``
	//   +   ``169.254.169.252/30``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#tunnel_inside_cidr Ec2VpnConnection#tunnel_inside_cidr}
	TunnelInsideCidr *string `field:"optional" json:"tunnelInsideCidr" yaml:"tunnelInsideCidr"`
	// The range of inside IPv6 addresses for the tunnel.
	//
	// Any specified CIDR blocks must be unique across all VPN connections that use the same transit gateway.
	//  Constraints: A size /126 CIDR block from the local ``fd00::/8`` range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_vpn_connection#tunnel_inside_ipv_6_cidr Ec2VpnConnection#tunnel_inside_ipv_6_cidr}
	TunnelInsideIpv6Cidr *string `field:"optional" json:"tunnelInsideIpv6Cidr" yaml:"tunnelInsideIpv6Cidr"`
}

