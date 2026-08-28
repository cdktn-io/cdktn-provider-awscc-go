// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice


type DevopsagentServiceServiceDetailsAzureIdentity struct {
	// Azure AD application client ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_service#client_id DevopsagentService#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Azure AD tenant ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_service#tenant_id DevopsagentService#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// ARN of the IAM role for web identity token exchange.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_service#web_identity_role_arn DevopsagentService#web_identity_role_arn}
	WebIdentityRoleArn *string `field:"optional" json:"webIdentityRoleArn" yaml:"webIdentityRoleArn"`
	// List of audiences for the web identity token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_service#web_identity_token_audiences DevopsagentService#web_identity_token_audiences}
	WebIdentityTokenAudiences *[]*string `field:"optional" json:"webIdentityTokenAudiences" yaml:"webIdentityTokenAudiences"`
}

