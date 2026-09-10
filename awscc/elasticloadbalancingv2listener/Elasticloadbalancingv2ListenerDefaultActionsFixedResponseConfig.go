// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2listener


type Elasticloadbalancingv2ListenerDefaultActionsFixedResponseConfig struct {
	// The content type.  Valid Values: text/plain | text/css | text/html | application/javascript | application/json.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticloadbalancingv2_listener#content_type Elasticloadbalancingv2Listener#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticloadbalancingv2_listener#message_body Elasticloadbalancingv2Listener#message_body}
	MessageBody *string `field:"optional" json:"messageBody" yaml:"messageBody"`
	// The HTTP response code (2XX, 4XX, or 5XX).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticloadbalancingv2_listener#status_code Elasticloadbalancingv2Listener#status_code}
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
}

