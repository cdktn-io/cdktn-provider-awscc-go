// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsconnection


type EventsConnectionAuthParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/events_connection#api_key_auth_parameters EventsConnection#api_key_auth_parameters}.
	ApiKeyAuthParameters *EventsConnectionAuthParametersApiKeyAuthParameters `field:"optional" json:"apiKeyAuthParameters" yaml:"apiKeyAuthParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/events_connection#basic_auth_parameters EventsConnection#basic_auth_parameters}.
	BasicAuthParameters *EventsConnectionAuthParametersBasicAuthParameters `field:"optional" json:"basicAuthParameters" yaml:"basicAuthParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/events_connection#connectivity_parameters EventsConnection#connectivity_parameters}.
	ConnectivityParameters *EventsConnectionAuthParametersConnectivityParameters `field:"optional" json:"connectivityParameters" yaml:"connectivityParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/events_connection#invocation_http_parameters EventsConnection#invocation_http_parameters}.
	InvocationHttpParameters *EventsConnectionAuthParametersInvocationHttpParameters `field:"optional" json:"invocationHttpParameters" yaml:"invocationHttpParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/events_connection#o_auth_parameters EventsConnection#o_auth_parameters}.
	OAuthParameters *EventsConnectionAuthParametersOAuthParameters `field:"optional" json:"oAuthParameters" yaml:"oAuthParameters"`
}

