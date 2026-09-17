// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package amplifybranch


type AmplifyBranchBasicAuthConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/amplify_branch#enable_basic_auth AmplifyBranch#enable_basic_auth}.
	EnableBasicAuth interface{} `field:"optional" json:"enableBasicAuth" yaml:"enableBasicAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/amplify_branch#password AmplifyBranch#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/amplify_branch#username AmplifyBranch#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

