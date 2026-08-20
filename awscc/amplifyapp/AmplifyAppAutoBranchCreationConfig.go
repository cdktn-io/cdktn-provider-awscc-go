// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package amplifyapp


type AmplifyAppAutoBranchCreationConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/amplify_app#auto_branch_creation_patterns AmplifyApp#auto_branch_creation_patterns}.
	AutoBranchCreationPatterns *[]*string `field:"optional" json:"autoBranchCreationPatterns" yaml:"autoBranchCreationPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/amplify_app#basic_auth_config AmplifyApp#basic_auth_config}.
	BasicAuthConfig *AmplifyAppAutoBranchCreationConfigBasicAuthConfig `field:"optional" json:"basicAuthConfig" yaml:"basicAuthConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/amplify_app#build_spec AmplifyApp#build_spec}.
	BuildSpec *string `field:"optional" json:"buildSpec" yaml:"buildSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/amplify_app#enable_auto_branch_creation AmplifyApp#enable_auto_branch_creation}.
	EnableAutoBranchCreation interface{} `field:"optional" json:"enableAutoBranchCreation" yaml:"enableAutoBranchCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/amplify_app#enable_auto_build AmplifyApp#enable_auto_build}.
	EnableAutoBuild interface{} `field:"optional" json:"enableAutoBuild" yaml:"enableAutoBuild"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/amplify_app#enable_performance_mode AmplifyApp#enable_performance_mode}.
	EnablePerformanceMode interface{} `field:"optional" json:"enablePerformanceMode" yaml:"enablePerformanceMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/amplify_app#enable_pull_request_preview AmplifyApp#enable_pull_request_preview}.
	EnablePullRequestPreview interface{} `field:"optional" json:"enablePullRequestPreview" yaml:"enablePullRequestPreview"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/amplify_app#environment_variables AmplifyApp#environment_variables}.
	EnvironmentVariables interface{} `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/amplify_app#framework AmplifyApp#framework}.
	Framework *string `field:"optional" json:"framework" yaml:"framework"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/amplify_app#pull_request_environment_name AmplifyApp#pull_request_environment_name}.
	PullRequestEnvironmentName *string `field:"optional" json:"pullRequestEnvironmentName" yaml:"pullRequestEnvironmentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/amplify_app#stage AmplifyApp#stage}.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

