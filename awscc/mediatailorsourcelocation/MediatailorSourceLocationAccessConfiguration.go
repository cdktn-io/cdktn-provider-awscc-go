// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorsourcelocation


type MediatailorSourceLocationAccessConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_source_location#access_type MediatailorSourceLocation#access_type}.
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
	// <p>AWS Secrets Manager access token configuration parameters.
	//
	// For information about Secrets Manager access token authentication, see <a href="https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-access-configuration-access-token.html">Working with AWS Secrets Manager access token authentication</a>.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediatailor_source_location#secrets_manager_access_token_configuration MediatailorSourceLocation#secrets_manager_access_token_configuration}
	SecretsManagerAccessTokenConfiguration *MediatailorSourceLocationAccessConfigurationSecretsManagerAccessTokenConfiguration `field:"optional" json:"secretsManagerAccessTokenConfiguration" yaml:"secretsManagerAccessTokenConfiguration"`
}

