// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDynatrace struct {
	// The API tokens used by Dynatrace API to authenticate various API calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appflow_connector_profile#api_token AppflowConnectorProfile#api_token}
	ApiToken *string `field:"optional" json:"apiToken" yaml:"apiToken"`
}

