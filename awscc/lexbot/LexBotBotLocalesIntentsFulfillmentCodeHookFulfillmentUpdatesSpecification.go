// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecification struct {
	// Determines whether fulfillment updates are sent to the user. When this field is true, updates are sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#active LexBot#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// Provides settings for a message that is sent to the user when a fulfillment Lambda function starts running.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#start_response LexBot#start_response}
	StartResponse *LexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponse `field:"optional" json:"startResponse" yaml:"startResponse"`
	// The length of time that the fulfillment Lambda function should run before it times out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#timeout_in_seconds LexBot#timeout_in_seconds}
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
	// Provides settings for a message that is sent periodically to the user while a fulfillment Lambda function is running.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#update_response LexBot#update_response}
	UpdateResponse *LexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationUpdateResponse `field:"optional" json:"updateResponse" yaml:"updateResponse"`
}

