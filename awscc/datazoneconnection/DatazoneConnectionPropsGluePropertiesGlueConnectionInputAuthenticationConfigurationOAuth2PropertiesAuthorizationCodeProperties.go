// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection


type DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOAuth2PropertiesAuthorizationCodeProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#authorization_code DatazoneConnection#authorization_code}.
	AuthorizationCode *string `field:"optional" json:"authorizationCode" yaml:"authorizationCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#redirect_uri DatazoneConnection#redirect_uri}.
	RedirectUri *string `field:"optional" json:"redirectUri" yaml:"redirectUri"`
}

