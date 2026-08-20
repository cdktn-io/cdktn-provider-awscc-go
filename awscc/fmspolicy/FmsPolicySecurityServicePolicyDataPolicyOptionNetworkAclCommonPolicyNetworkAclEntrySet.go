// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fmspolicy


type FmsPolicySecurityServicePolicyDataPolicyOptionNetworkAclCommonPolicyNetworkAclEntrySet struct {
	// NetworkAcl entry list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/fms_policy#first_entries FmsPolicy#first_entries}
	FirstEntries interface{} `field:"optional" json:"firstEntries" yaml:"firstEntries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/fms_policy#force_remediate_for_first_entries FmsPolicy#force_remediate_for_first_entries}.
	ForceRemediateForFirstEntries interface{} `field:"optional" json:"forceRemediateForFirstEntries" yaml:"forceRemediateForFirstEntries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/fms_policy#force_remediate_for_last_entries FmsPolicy#force_remediate_for_last_entries}.
	ForceRemediateForLastEntries interface{} `field:"optional" json:"forceRemediateForLastEntries" yaml:"forceRemediateForLastEntries"`
	// NetworkAcl entry list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/fms_policy#last_entries FmsPolicy#last_entries}
	LastEntries interface{} `field:"optional" json:"lastEntries" yaml:"lastEntries"`
}

