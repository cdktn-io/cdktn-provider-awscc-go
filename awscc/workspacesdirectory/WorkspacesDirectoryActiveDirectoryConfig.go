// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory


type WorkspacesDirectoryActiveDirectoryConfig struct {
	// The name of the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#domain_name WorkspacesDirectory#domain_name}
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Indicates the secret ARN on the service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#service_account_secret_arn WorkspacesDirectory#service_account_secret_arn}
	ServiceAccountSecretArn *string `field:"optional" json:"serviceAccountSecretArn" yaml:"serviceAccountSecretArn"`
}

