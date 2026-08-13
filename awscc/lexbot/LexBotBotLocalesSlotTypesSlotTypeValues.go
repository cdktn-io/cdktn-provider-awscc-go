// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesSlotTypesSlotTypeValues struct {
	// Defines one of the values for a slot type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lex_bot#sample_value LexBot#sample_value}
	SampleValue *LexBotBotLocalesSlotTypesSlotTypeValuesSampleValue `field:"optional" json:"sampleValue" yaml:"sampleValue"`
	// Additional values related to the slot type entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lex_bot#synonyms LexBot#synonyms}
	Synonyms interface{} `field:"optional" json:"synonyms" yaml:"synonyms"`
}

