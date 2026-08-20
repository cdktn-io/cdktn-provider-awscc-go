// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowTags struct {
	// A string used to identify this tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appflow_flow#key AppflowFlow#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A string containing the value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appflow_flow#value AppflowFlow#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

