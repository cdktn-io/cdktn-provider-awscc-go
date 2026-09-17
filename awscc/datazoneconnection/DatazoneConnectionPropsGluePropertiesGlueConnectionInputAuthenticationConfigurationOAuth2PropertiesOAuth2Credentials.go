// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection


type DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2Credentials struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#access_token DatazoneConnection#access_token}.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#jwt_token DatazoneConnection#jwt_token}.
	JwtToken *string `field:"optional" json:"jwtToken" yaml:"jwtToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#refresh_token DatazoneConnection#refresh_token}.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#user_managed_client_application_client_secret DatazoneConnection#user_managed_client_application_client_secret}.
	UserManagedClientApplicationClientSecret *string `field:"optional" json:"userManagedClientApplicationClientSecret" yaml:"userManagedClientApplicationClientSecret"`
}

