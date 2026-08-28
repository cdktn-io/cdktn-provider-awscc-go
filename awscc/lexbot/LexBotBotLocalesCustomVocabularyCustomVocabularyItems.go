// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesCustomVocabularyCustomVocabularyItems struct {
	// Phrase that should be recognized.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#phrase LexBot#phrase}
	Phrase *string `field:"optional" json:"phrase" yaml:"phrase"`
	// The degree to which the phrase recognition is boosted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#weight LexBot#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

