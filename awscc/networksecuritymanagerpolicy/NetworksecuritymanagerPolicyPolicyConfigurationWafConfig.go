// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecuritymanagerpolicy


type NetworksecuritymanagerPolicyPolicyConfigurationWafConfig struct {
	// Conflict-resolution strategy applied to AWS WAF policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/networksecuritymanager_policy#conflict_resolution NetworksecuritymanagerPolicy#conflict_resolution}
	ConflictResolution *string `field:"optional" json:"conflictResolution" yaml:"conflictResolution"`
	// Controls how Network Security Manager handles remediation when a resource already has a customer-created WebACL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/networksecuritymanager_policy#existing_customer_web_acl_resolution NetworksecuritymanagerPolicy#existing_customer_web_acl_resolution}
	ExistingCustomerWebAclResolution *string `field:"optional" json:"existingCustomerWebAclResolution" yaml:"existingCustomerWebAclResolution"`
}

