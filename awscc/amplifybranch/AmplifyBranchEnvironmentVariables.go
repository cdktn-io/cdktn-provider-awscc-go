// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package amplifybranch


type AmplifyBranchEnvironmentVariables struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/amplify_branch#name AmplifyBranch#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/amplify_branch#value AmplifyBranch#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

