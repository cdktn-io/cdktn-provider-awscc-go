// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorCustom struct {
	// A map for properties for custom authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appflow_connector_profile#credentials_map AppflowConnectorProfile#credentials_map}
	CredentialsMap *map[string]*string `field:"optional" json:"credentialsMap" yaml:"credentialsMap"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appflow_connector_profile#custom_authentication_type AppflowConnectorProfile#custom_authentication_type}.
	CustomAuthenticationType *string `field:"optional" json:"customAuthenticationType" yaml:"customAuthenticationType"`
}

