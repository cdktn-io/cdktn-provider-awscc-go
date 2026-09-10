// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubconfigurationpolicy


type SecurityhubConfigurationPolicyConfigurationPolicySecurityHub struct {
	// A list that defines which security standards are enabled in the configuration policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_configuration_policy#enabled_standard_identifiers SecurityhubConfigurationPolicy#enabled_standard_identifiers}
	EnabledStandardIdentifiers *[]*string `field:"optional" json:"enabledStandardIdentifiers" yaml:"enabledStandardIdentifiers"`
	// An object that defines which security controls are enabled in an AWS Security Hub configuration policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_configuration_policy#security_controls_configuration SecurityhubConfigurationPolicy#security_controls_configuration}
	SecurityControlsConfiguration *SecurityhubConfigurationPolicyConfigurationPolicySecurityHubSecurityControlsConfiguration `field:"optional" json:"securityControlsConfiguration" yaml:"securityControlsConfiguration"`
	// Indicates whether Security Hub is enabled in the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_configuration_policy#service_enabled SecurityhubConfigurationPolicy#service_enabled}
	ServiceEnabled interface{} `field:"optional" json:"serviceEnabled" yaml:"serviceEnabled"`
}

