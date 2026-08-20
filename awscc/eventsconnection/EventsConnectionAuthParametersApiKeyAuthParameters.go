// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsconnection


type EventsConnectionAuthParametersApiKeyAuthParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/events_connection#api_key_name EventsConnection#api_key_name}.
	ApiKeyName *string `field:"optional" json:"apiKeyName" yaml:"apiKeyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/events_connection#api_key_value EventsConnection#api_key_value}.
	ApiKeyValue *string `field:"optional" json:"apiKeyValue" yaml:"apiKeyValue"`
}

