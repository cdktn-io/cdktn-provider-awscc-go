// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsconnection


type EventsConnectionAuthParametersOAuthParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_connection#authorization_endpoint EventsConnection#authorization_endpoint}.
	AuthorizationEndpoint *string `field:"optional" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_connection#client_parameters EventsConnection#client_parameters}.
	ClientParameters *EventsConnectionAuthParametersOAuthParametersClientParameters `field:"optional" json:"clientParameters" yaml:"clientParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_connection#http_method EventsConnection#http_method}.
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_connection#o_auth_http_parameters EventsConnection#o_auth_http_parameters}.
	OAuthHttpParameters *EventsConnectionAuthParametersOAuthParametersOAuthHttpParameters `field:"optional" json:"oAuthHttpParameters" yaml:"oAuthHttpParameters"`
}

