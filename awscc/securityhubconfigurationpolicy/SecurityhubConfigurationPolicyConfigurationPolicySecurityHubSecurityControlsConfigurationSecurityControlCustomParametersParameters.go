// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubconfigurationpolicy


type SecurityhubConfigurationPolicyConfigurationPolicySecurityHubSecurityControlsConfigurationSecurityControlCustomParametersParameters struct {
	// An object that includes the data type of a security control parameter and its current value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/securityhub_configuration_policy#value SecurityhubConfigurationPolicy#value}
	Value *SecurityhubConfigurationPolicyConfigurationPolicySecurityHubSecurityControlsConfigurationSecurityControlCustomParametersParametersValue `field:"optional" json:"value" yaml:"value"`
	// Identifies whether a control parameter uses a custom user-defined value or subscribes to the default AWS Security Hub behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/securityhub_configuration_policy#value_type SecurityhubConfigurationPolicy#value_type}
	ValueType *string `field:"optional" json:"valueType" yaml:"valueType"`
}

