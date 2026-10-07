// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber


type Eventsv2SubscriberInvokeConfigurationUniversalTargetParameters struct {
	// JSON string or JSONata expression that produces the API request.
	//
	// Supports {% ... %} JSONata expressions for dynamic values from the event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#input Eventsv2Subscriber#input}
	Input *string `field:"optional" json:"input" yaml:"input"`
	// Timeout in seconds for each invocation of the target (1-30, default 30).
	//
	// Must be a literal integer written as a string; JSONata expressions are not supported for this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#invocation_timeout_seconds Eventsv2Subscriber#invocation_timeout_seconds}
	InvocationTimeoutSeconds *string `field:"optional" json:"invocationTimeoutSeconds" yaml:"invocationTimeoutSeconds"`
}

