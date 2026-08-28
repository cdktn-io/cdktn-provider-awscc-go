// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentapplication


type SecurityagentApplicationTags struct {
	// The key name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityagent_application#key SecurityagentApplication#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityagent_application#value SecurityagentApplication#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

