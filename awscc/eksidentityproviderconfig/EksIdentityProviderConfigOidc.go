// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eksidentityproviderconfig


type EksIdentityProviderConfigOidc struct {
	// This is also known as audience.
	//
	// The ID for the client application that makes authentication requests to the OpenID identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/eks_identity_provider_config#client_id EksIdentityProviderConfig#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The JWT claim that the provider uses to return your groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/eks_identity_provider_config#groups_claim EksIdentityProviderConfig#groups_claim}
	GroupsClaim *string `field:"optional" json:"groupsClaim" yaml:"groupsClaim"`
	// The prefix that is prepended to group claims to prevent clashes with existing names (such as system: groups).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/eks_identity_provider_config#groups_prefix EksIdentityProviderConfig#groups_prefix}
	GroupsPrefix *string `field:"optional" json:"groupsPrefix" yaml:"groupsPrefix"`
	// The URL of the OpenID identity provider that allows the API server to discover public signing keys for verifying tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/eks_identity_provider_config#issuer_url EksIdentityProviderConfig#issuer_url}
	IssuerUrl *string `field:"optional" json:"issuerUrl" yaml:"issuerUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/eks_identity_provider_config#required_claims EksIdentityProviderConfig#required_claims}.
	RequiredClaims interface{} `field:"optional" json:"requiredClaims" yaml:"requiredClaims"`
	// The JSON Web Token (JWT) claim to use as the username.
	//
	// The default is sub, which is expected to be a unique identifier of the end user. You can choose other claims, such as email or name, depending on the OpenID identity provider. Claims other than email are prefixed with the issuer URL to prevent naming clashes with other plug-ins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/eks_identity_provider_config#username_claim EksIdentityProviderConfig#username_claim}
	UsernameClaim *string `field:"optional" json:"usernameClaim" yaml:"usernameClaim"`
	// The prefix that is prepended to username claims to prevent clashes with existing names.
	//
	// If you do not provide this field, and username is a value other than email, the prefix defaults to issuerurl#. You can use the value - to disable all prefixing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/eks_identity_provider_config#username_prefix EksIdentityProviderConfig#username_prefix}
	UsernamePrefix *string `field:"optional" json:"usernamePrefix" yaml:"usernamePrefix"`
}

