// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appfabricappauthorization


type AppfabricAppAuthorizationCredentialOauth2Credential struct {
	// The client ID of the client application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/appfabric_app_authorization#client_id AppfabricAppAuthorization#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The client secret of the client application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/appfabric_app_authorization#client_secret AppfabricAppAuthorization#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
}

