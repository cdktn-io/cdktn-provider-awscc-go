// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntents struct {
	// A description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#description LexBot#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Settings that determine the Lambda function that Amazon Lex uses for processing user responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#dialog_code_hook LexBot#dialog_code_hook}
	DialogCodeHook *LexBotBotLocalesIntentsDialogCodeHook `field:"optional" json:"dialogCodeHook" yaml:"dialogCodeHook"`
	// Settings that determine if a Lambda function should be invoked to fulfill a specific intent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#fulfillment_code_hook LexBot#fulfillment_code_hook}
	FulfillmentCodeHook *LexBotBotLocalesIntentsFulfillmentCodeHook `field:"optional" json:"fulfillmentCodeHook" yaml:"fulfillmentCodeHook"`
	// The list of input contexts specified for the intent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#input_contexts LexBot#input_contexts}
	InputContexts interface{} `field:"optional" json:"inputContexts" yaml:"inputContexts"`
	// Response that Amazon Lex sends to the user when the intent is closed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#intent_closing_setting LexBot#intent_closing_setting}
	IntentClosingSetting *LexBotBotLocalesIntentsIntentClosingSetting `field:"optional" json:"intentClosingSetting" yaml:"intentClosingSetting"`
	// Prompts that Amazon Lex sends to the user to confirm the completion of an intent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#intent_confirmation_setting LexBot#intent_confirmation_setting}
	IntentConfirmationSetting *LexBotBotLocalesIntentsIntentConfirmationSetting `field:"optional" json:"intentConfirmationSetting" yaml:"intentConfirmationSetting"`
	// Configuration for searching a Amazon Kendra index specified for the intent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#kendra_configuration LexBot#kendra_configuration}
	KendraConfiguration *LexBotBotLocalesIntentsKendraConfiguration `field:"optional" json:"kendraConfiguration" yaml:"kendraConfiguration"`
	// Unique name for a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#name LexBot#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of contexts that the intent activates when it is fulfilled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#output_contexts LexBot#output_contexts}
	OutputContexts interface{} `field:"optional" json:"outputContexts" yaml:"outputContexts"`
	// A unique identifier for the built-in intent to base this intent on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#parent_intent_signature LexBot#parent_intent_signature}
	ParentIntentSignature *string `field:"optional" json:"parentIntentSignature" yaml:"parentIntentSignature"`
	// An array of sample utterances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#sample_utterances LexBot#sample_utterances}
	SampleUtterances interface{} `field:"optional" json:"sampleUtterances" yaml:"sampleUtterances"`
	// List for slot priorities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#slot_priorities LexBot#slot_priorities}
	SlotPriorities interface{} `field:"optional" json:"slotPriorities" yaml:"slotPriorities"`
	// List of slots.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#slots LexBot#slots}
	Slots interface{} `field:"optional" json:"slots" yaml:"slots"`
}

