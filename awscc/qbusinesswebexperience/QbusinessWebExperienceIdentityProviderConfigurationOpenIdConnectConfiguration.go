// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinesswebexperience


type QbusinessWebExperienceIdentityProviderConfigurationOpenIdConnectConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/qbusiness_web_experience#secrets_arn QbusinessWebExperience#secrets_arn}.
	SecretsArn *string `field:"optional" json:"secretsArn" yaml:"secretsArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/qbusiness_web_experience#secrets_role QbusinessWebExperience#secrets_role}.
	SecretsRole *string `field:"optional" json:"secretsRole" yaml:"secretsRole"`
}

