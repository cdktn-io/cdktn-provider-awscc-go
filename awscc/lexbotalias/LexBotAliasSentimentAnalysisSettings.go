// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbotalias


type LexBotAliasSentimentAnalysisSettings struct {
	// Enable to call Amazon Comprehend for Sentiment natively within Lex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot_alias#detect_sentiment LexBotAlias#detect_sentiment}
	DetectSentiment interface{} `field:"optional" json:"detectSentiment" yaml:"detectSentiment"`
}

