// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotTestBotAliasSettingsSentimentAnalysisSettings struct {
	// Enable to call Amazon Comprehend for Sentiment natively within Lex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#detect_sentiment LexBot#detect_sentiment}
	DetectSentiment interface{} `field:"optional" json:"detectSentiment" yaml:"detectSentiment"`
}

