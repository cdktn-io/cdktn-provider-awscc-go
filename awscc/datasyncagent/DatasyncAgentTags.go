// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasyncagent


type DatasyncAgentTags struct {
	// The key for an AWS resource tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datasync_agent#key DatasyncAgent#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for an AWS resource tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datasync_agent#value DatasyncAgent#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

