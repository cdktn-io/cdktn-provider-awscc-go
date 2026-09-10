// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchservicedomain


type OpensearchserviceDomainAdvancedSecurityOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/opensearchservice_domain#anonymous_auth_enabled OpensearchserviceDomain#anonymous_auth_enabled}.
	AnonymousAuthEnabled interface{} `field:"optional" json:"anonymousAuthEnabled" yaml:"anonymousAuthEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/opensearchservice_domain#enabled OpensearchserviceDomain#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/opensearchservice_domain#iam_federation_options OpensearchserviceDomain#iam_federation_options}.
	IamFederationOptions *OpensearchserviceDomainAdvancedSecurityOptionsIamFederationOptions `field:"optional" json:"iamFederationOptions" yaml:"iamFederationOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/opensearchservice_domain#internal_user_database_enabled OpensearchserviceDomain#internal_user_database_enabled}.
	InternalUserDatabaseEnabled interface{} `field:"optional" json:"internalUserDatabaseEnabled" yaml:"internalUserDatabaseEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/opensearchservice_domain#jwt_options OpensearchserviceDomain#jwt_options}.
	JwtOptions *OpensearchserviceDomainAdvancedSecurityOptionsJwtOptions `field:"optional" json:"jwtOptions" yaml:"jwtOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/opensearchservice_domain#master_user_options OpensearchserviceDomain#master_user_options}.
	MasterUserOptions *OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptions `field:"optional" json:"masterUserOptions" yaml:"masterUserOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/opensearchservice_domain#saml_options OpensearchserviceDomain#saml_options}.
	SamlOptions *OpensearchserviceDomainAdvancedSecurityOptionsSamlOptions `field:"optional" json:"samlOptions" yaml:"samlOptions"`
}

