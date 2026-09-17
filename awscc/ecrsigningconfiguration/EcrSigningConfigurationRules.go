// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecrsigningconfiguration


type EcrSigningConfigurationRules struct {
	// AWS Signer signing profile ARN to use for matched repositories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ecr_signing_configuration#signing_profile_arn EcrSigningConfiguration#signing_profile_arn}
	SigningProfileArn *string `field:"required" json:"signingProfileArn" yaml:"signingProfileArn"`
	// Optional array of repository filters.
	//
	// If omitted, the rule matches all repositories. If provided, must contain at least one filter. Empty arrays are not allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ecr_signing_configuration#repository_filters EcrSigningConfiguration#repository_filters}
	RepositoryFilters interface{} `field:"optional" json:"repositoryFilters" yaml:"repositoryFilters"`
}

