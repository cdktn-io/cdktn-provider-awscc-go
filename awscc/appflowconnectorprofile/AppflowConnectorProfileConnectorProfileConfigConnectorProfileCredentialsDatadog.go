// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDatadog struct {
	// A unique alphanumeric identi?er used to authenticate a user, developer, or calling program to your API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appflow_connector_profile#api_key AppflowConnectorProfile#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Application keys, in conjunction with your API key, give you full access to Datadog?s programmatic API. Application keys are associated with the user account that created them. The application key is used to log all requests made to the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appflow_connector_profile#application_key AppflowConnectorProfile#application_key}
	ApplicationKey *string `field:"optional" json:"applicationKey" yaml:"applicationKey"`
}

