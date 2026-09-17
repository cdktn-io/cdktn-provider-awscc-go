// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fmspolicy


type FmsPolicySecurityServicePolicyData struct {
	// Firewall policy type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/fms_policy#type FmsPolicy#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Firewall managed service data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/fms_policy#managed_service_data FmsPolicy#managed_service_data}
	ManagedServiceData *string `field:"optional" json:"managedServiceData" yaml:"managedServiceData"`
	// Firewall policy option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/fms_policy#policy_option FmsPolicy#policy_option}
	PolicyOption *FmsPolicySecurityServicePolicyDataPolicyOption `field:"optional" json:"policyOption" yaml:"policyOption"`
}

