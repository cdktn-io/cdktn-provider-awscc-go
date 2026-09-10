// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmpatchbaseline


type SsmPatchBaselineApprovalRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssm_patch_baseline#patch_rules SsmPatchBaseline#patch_rules}.
	PatchRules interface{} `field:"optional" json:"patchRules" yaml:"patchRules"`
}

