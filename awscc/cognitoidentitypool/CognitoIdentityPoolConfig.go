// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitoidentitypool

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CognitoIdentityPoolConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_identity_pool#allow_unauthenticated_identities CognitoIdentityPool#allow_unauthenticated_identities}.
	AllowUnauthenticatedIdentities interface{} `field:"required" json:"allowUnauthenticatedIdentities" yaml:"allowUnauthenticatedIdentities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_identity_pool#allow_classic_flow CognitoIdentityPool#allow_classic_flow}.
	AllowClassicFlow interface{} `field:"optional" json:"allowClassicFlow" yaml:"allowClassicFlow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_identity_pool#cognito_events CognitoIdentityPool#cognito_events}.
	CognitoEvents *string `field:"optional" json:"cognitoEvents" yaml:"cognitoEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_identity_pool#cognito_identity_providers CognitoIdentityPool#cognito_identity_providers}.
	CognitoIdentityProviders interface{} `field:"optional" json:"cognitoIdentityProviders" yaml:"cognitoIdentityProviders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_identity_pool#cognito_streams CognitoIdentityPool#cognito_streams}.
	CognitoStreams *CognitoIdentityPoolCognitoStreams `field:"optional" json:"cognitoStreams" yaml:"cognitoStreams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_identity_pool#developer_provider_name CognitoIdentityPool#developer_provider_name}.
	DeveloperProviderName *string `field:"optional" json:"developerProviderName" yaml:"developerProviderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_identity_pool#identity_pool_name CognitoIdentityPool#identity_pool_name}.
	IdentityPoolName *string `field:"optional" json:"identityPoolName" yaml:"identityPoolName"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_identity_pool#identity_pool_tags CognitoIdentityPool#identity_pool_tags}
	IdentityPoolTags interface{} `field:"optional" json:"identityPoolTags" yaml:"identityPoolTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_identity_pool#open_id_connect_provider_ar_ns CognitoIdentityPool#open_id_connect_provider_ar_ns}.
	OpenIdConnectProviderArNs *[]*string `field:"optional" json:"openIdConnectProviderArNs" yaml:"openIdConnectProviderArNs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_identity_pool#push_sync CognitoIdentityPool#push_sync}.
	PushSync *CognitoIdentityPoolPushSync `field:"optional" json:"pushSync" yaml:"pushSync"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_identity_pool#saml_provider_ar_ns CognitoIdentityPool#saml_provider_ar_ns}.
	SamlProviderArNs *[]*string `field:"optional" json:"samlProviderArNs" yaml:"samlProviderArNs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cognito_identity_pool#supported_login_providers CognitoIdentityPool#supported_login_providers}.
	SupportedLoginProviders *string `field:"optional" json:"supportedLoginProviders" yaml:"supportedLoginProviders"`
}

