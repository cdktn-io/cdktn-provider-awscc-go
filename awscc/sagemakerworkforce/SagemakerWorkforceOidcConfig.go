// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkforce


type SagemakerWorkforceOidcConfig struct {
	// A string to string map of identifiers specific to the custom identity provider (IdP) being used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_workforce#authentication_request_extra_params SagemakerWorkforce#authentication_request_extra_params}
	AuthenticationRequestExtraParams *map[string]*string `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The OIDC IdP authorization endpoint used to configure your private workforce.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_workforce#authorization_endpoint SagemakerWorkforce#authorization_endpoint}
	AuthorizationEndpoint *string `field:"optional" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The OIDC IdP client ID used to configure your private workforce.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_workforce#client_id SagemakerWorkforce#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The OIDC IdP client secret used to configure your private workforce.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_workforce#client_secret SagemakerWorkforce#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The OIDC IdP issuer used to configure your private workforce.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_workforce#issuer SagemakerWorkforce#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// The OIDC IdP JSON Web Key Set (Jwks) URI used to configure your private workforce.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_workforce#jwks_uri SagemakerWorkforce#jwks_uri}
	JwksUri *string `field:"optional" json:"jwksUri" yaml:"jwksUri"`
	// The OIDC IdP logout endpoint used to configure your private workforce.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_workforce#logout_endpoint SagemakerWorkforce#logout_endpoint}
	LogoutEndpoint *string `field:"optional" json:"logoutEndpoint" yaml:"logoutEndpoint"`
	// An array of string identifiers used to refer to the specific pieces of user data or claims that the client application wants to access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_workforce#scope SagemakerWorkforce#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The OIDC IdP token endpoint used to configure your private workforce.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_workforce#token_endpoint SagemakerWorkforce#token_endpoint}
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// The OIDC IdP user info endpoint used to configure your private workforce.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_workforce#user_info_endpoint SagemakerWorkforce#user_info_endpoint}
	UserInfoEndpoint *string `field:"optional" json:"userInfoEndpoint" yaml:"userInfoEndpoint"`
}

