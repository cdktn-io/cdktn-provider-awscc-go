// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice


type DevopsagentServiceServiceDetailsGitLab struct {
	// Optional GitLab group ID for group-level access tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/devopsagent_service#group_id DevopsagentService#group_id}
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// GitLab instance URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/devopsagent_service#target_url DevopsagentService#target_url}
	TargetUrl *string `field:"optional" json:"targetUrl" yaml:"targetUrl"`
	// Type of GitLab access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/devopsagent_service#token_type DevopsagentService#token_type}
	TokenType *string `field:"optional" json:"tokenType" yaml:"tokenType"`
	// GitLab access token value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/devopsagent_service#token_value DevopsagentService#token_value}
	TokenValue *string `field:"optional" json:"tokenValue" yaml:"tokenValue"`
}

