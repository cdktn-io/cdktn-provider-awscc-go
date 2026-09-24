// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessSkills struct {
	// The filesystem path to the skill definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_harness#path BedrockagentcoreHarness#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

