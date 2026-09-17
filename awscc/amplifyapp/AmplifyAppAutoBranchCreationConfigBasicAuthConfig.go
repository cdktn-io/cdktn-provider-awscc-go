// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package amplifyapp


type AmplifyAppAutoBranchCreationConfigBasicAuthConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/amplify_app#enable_basic_auth AmplifyApp#enable_basic_auth}.
	EnableBasicAuth interface{} `field:"optional" json:"enableBasicAuth" yaml:"enableBasicAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/amplify_app#password AmplifyApp#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/amplify_app#username AmplifyApp#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

