// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fmspolicy


type FmsPolicySecurityServicePolicyDataPolicyOptionNetworkAclCommonPolicyNetworkAclEntrySetLastEntriesPortRange struct {
	// From Port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/fms_policy#from FmsPolicy#from}
	From *float64 `field:"optional" json:"from" yaml:"from"`
	// To Port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/fms_policy#to FmsPolicy#to}
	To *float64 `field:"optional" json:"to" yaml:"to"`
}

