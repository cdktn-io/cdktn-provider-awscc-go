// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubconfigurationpolicy


type SecurityhubConfigurationPolicyConfigurationPolicySecurityHubSecurityControlsConfigurationSecurityControlCustomParameters struct {
	// An object that specifies parameter values for a control in a configuration policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_configuration_policy#parameters SecurityhubConfigurationPolicy#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The ID of the security control.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_configuration_policy#security_control_id SecurityhubConfigurationPolicy#security_control_id}
	SecurityControlId *string `field:"optional" json:"securityControlId" yaml:"securityControlId"`
}

