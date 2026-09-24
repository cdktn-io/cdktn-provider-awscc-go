// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectsecurityprofile


type ConnectSecurityProfileGranularAccessControlConfigurationDataTableAccessControlConfiguration struct {
	// Contains the configuration for record-based access control.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_security_profile#primary_attribute_access_control_configuration ConnectSecurityProfile#primary_attribute_access_control_configuration}
	PrimaryAttributeAccessControlConfiguration *ConnectSecurityProfileGranularAccessControlConfigurationDataTableAccessControlConfigurationPrimaryAttributeAccessControlConfiguration `field:"optional" json:"primaryAttributeAccessControlConfiguration" yaml:"primaryAttributeAccessControlConfiguration"`
}

