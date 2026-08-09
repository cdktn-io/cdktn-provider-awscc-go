// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package awsexternalanthropicworkspace


type AwsexternalanthropicWorkspaceTags struct {
	// The key name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/awsexternalanthropic_workspace#key AwsexternalanthropicWorkspace#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/awsexternalanthropic_workspace#value AwsexternalanthropicWorkspace#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

