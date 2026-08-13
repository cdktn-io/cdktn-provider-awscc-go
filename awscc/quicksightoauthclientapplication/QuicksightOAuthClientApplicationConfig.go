// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightoauthclientapplication

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightOAuthClientApplicationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_o_auth_client_application#name QuicksightOAuthClientApplication#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_o_auth_client_application#o_auth_client_application_id QuicksightOAuthClientApplication#o_auth_client_application_id}.
	OAuthClientApplicationId *string `field:"required" json:"oAuthClientApplicationId" yaml:"oAuthClientApplicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_o_auth_client_application#o_auth_client_authentication_type QuicksightOAuthClientApplication#o_auth_client_authentication_type}.
	OAuthClientAuthenticationType *string `field:"required" json:"oAuthClientAuthenticationType" yaml:"oAuthClientAuthenticationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_o_auth_client_application#o_auth_token_endpoint_url QuicksightOAuthClientApplication#o_auth_token_endpoint_url}.
	OAuthTokenEndpointUrl *string `field:"required" json:"oAuthTokenEndpointUrl" yaml:"oAuthTokenEndpointUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_o_auth_client_application#client_id QuicksightOAuthClientApplication#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_o_auth_client_application#client_secret QuicksightOAuthClientApplication#client_secret}.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_o_auth_client_application#data_source_type QuicksightOAuthClientApplication#data_source_type}.
	DataSourceType *string `field:"optional" json:"dataSourceType" yaml:"dataSourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_o_auth_client_application#identity_provider_vpc_connection_properties QuicksightOAuthClientApplication#identity_provider_vpc_connection_properties}.
	IdentityProviderVpcConnectionProperties *QuicksightOAuthClientApplicationIdentityProviderVpcConnectionProperties `field:"optional" json:"identityProviderVpcConnectionProperties" yaml:"identityProviderVpcConnectionProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_o_auth_client_application#o_auth_authorization_endpoint_url QuicksightOAuthClientApplication#o_auth_authorization_endpoint_url}.
	OAuthAuthorizationEndpointUrl *string `field:"optional" json:"oAuthAuthorizationEndpointUrl" yaml:"oAuthAuthorizationEndpointUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_o_auth_client_application#o_auth_scopes QuicksightOAuthClientApplication#o_auth_scopes}.
	OAuthScopes *string `field:"optional" json:"oAuthScopes" yaml:"oAuthScopes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_o_auth_client_application#tags QuicksightOAuthClientApplication#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

