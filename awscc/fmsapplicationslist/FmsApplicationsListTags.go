// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fmsapplicationslist


type FmsApplicationsListTags struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/fms_applications_list#key FmsApplicationsList#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/fms_applications_list#value FmsApplicationsList#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

