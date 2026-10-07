// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory


type WorkspacesDirectoryMicrosoftEntraConfig struct {
	// The Amazon Resource Name (ARN) of the application config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#application_config_secret_arn WorkspacesDirectory#application_config_secret_arn}
	ApplicationConfigSecretArn *string `field:"optional" json:"applicationConfigSecretArn" yaml:"applicationConfigSecretArn"`
	// The identifier of the tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#tenant_id WorkspacesDirectory#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

