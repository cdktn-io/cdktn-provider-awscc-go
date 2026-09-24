// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecification struct {
	// A list of message groups that Amazon Lex uses to respond the user input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lex_bot#failure_response LexBot#failure_response}
	FailureResponse *LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponse `field:"optional" json:"failureResponse" yaml:"failureResponse"`
	// A list of message groups that Amazon Lex uses to respond the user input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lex_bot#success_response LexBot#success_response}
	SuccessResponse *LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponse `field:"optional" json:"successResponse" yaml:"successResponse"`
	// A list of message groups that Amazon Lex uses to respond the user input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lex_bot#timeout_response LexBot#timeout_response}
	TimeoutResponse *LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutResponse `field:"optional" json:"timeoutResponse" yaml:"timeoutResponse"`
}

