// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubconfigurationpolicy


type SecurityhubConfigurationPolicyConfigurationPolicySecurityHubSecurityControlsConfiguration struct {
	// A list of security controls that are disabled in the configuration policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/securityhub_configuration_policy#disabled_security_control_identifiers SecurityhubConfigurationPolicy#disabled_security_control_identifiers}
	DisabledSecurityControlIdentifiers *[]*string `field:"optional" json:"disabledSecurityControlIdentifiers" yaml:"disabledSecurityControlIdentifiers"`
	// A list of security controls that are enabled in the configuration policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/securityhub_configuration_policy#enabled_security_control_identifiers SecurityhubConfigurationPolicy#enabled_security_control_identifiers}
	EnabledSecurityControlIdentifiers *[]*string `field:"optional" json:"enabledSecurityControlIdentifiers" yaml:"enabledSecurityControlIdentifiers"`
	// A list of security controls and control parameter values that are included in a configuration policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/securityhub_configuration_policy#security_control_custom_parameters SecurityhubConfigurationPolicy#security_control_custom_parameters}
	SecurityControlCustomParameters interface{} `field:"optional" json:"securityControlCustomParameters" yaml:"securityControlCustomParameters"`
}

