// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appfabricappauthorization


type AppfabricAppAuthorizationCredential struct {
	// Contains API key credential information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/appfabric_app_authorization#api_key_credential AppfabricAppAuthorization#api_key_credential}
	ApiKeyCredential *AppfabricAppAuthorizationCredentialApiKeyCredential `field:"optional" json:"apiKeyCredential" yaml:"apiKeyCredential"`
	// Contains OAuth2 client credential information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/appfabric_app_authorization#oauth_2_credential AppfabricAppAuthorization#oauth_2_credential}
	Oauth2Credential *AppfabricAppAuthorizationCredentialOauth2Credential `field:"optional" json:"oauth2Credential" yaml:"oauth2Credential"`
}

