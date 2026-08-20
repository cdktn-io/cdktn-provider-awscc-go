// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkfirewalltlsinspectionconfiguration


type NetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsCheckCertificateRevocationStatus struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkfirewall_tls_inspection_configuration#revoked_status_action NetworkfirewallTlsInspectionConfiguration#revoked_status_action}.
	RevokedStatusAction *string `field:"optional" json:"revokedStatusAction" yaml:"revokedStatusAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/networkfirewall_tls_inspection_configuration#unknown_status_action NetworkfirewallTlsInspectionConfiguration#unknown_status_action}.
	UnknownStatusAction *string `field:"optional" json:"unknownStatusAction" yaml:"unknownStatusAction"`
}

