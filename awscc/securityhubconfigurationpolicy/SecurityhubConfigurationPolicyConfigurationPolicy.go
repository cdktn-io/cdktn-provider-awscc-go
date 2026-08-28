// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubconfigurationpolicy


type SecurityhubConfigurationPolicyConfigurationPolicy struct {
	// An object that defines how AWS Security Hub is configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityhub_configuration_policy#security_hub SecurityhubConfigurationPolicy#security_hub}
	SecurityHub *SecurityhubConfigurationPolicyConfigurationPolicySecurityHub `field:"optional" json:"securityHub" yaml:"securityHub"`
}

