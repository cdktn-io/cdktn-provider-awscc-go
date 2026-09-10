// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectsecurityprofile


type ConnectSecurityProfileGranularAccessControlConfiguration struct {
	// Defines the access control configuration for data tables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_security_profile#data_table_access_control_configuration ConnectSecurityProfile#data_table_access_control_configuration}
	DataTableAccessControlConfiguration *ConnectSecurityProfileGranularAccessControlConfigurationDataTableAccessControlConfiguration `field:"optional" json:"dataTableAccessControlConfiguration" yaml:"dataTableAccessControlConfiguration"`
}

