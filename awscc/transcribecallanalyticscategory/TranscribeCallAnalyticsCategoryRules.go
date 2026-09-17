// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transcribecallanalyticscategory


type TranscribeCallAnalyticsCategoryRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/transcribe_call_analytics_category#interruption_filter TranscribeCallAnalyticsCategory#interruption_filter}.
	InterruptionFilter *TranscribeCallAnalyticsCategoryRulesInterruptionFilter `field:"optional" json:"interruptionFilter" yaml:"interruptionFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/transcribe_call_analytics_category#non_talk_time_filter TranscribeCallAnalyticsCategory#non_talk_time_filter}.
	NonTalkTimeFilter *TranscribeCallAnalyticsCategoryRulesNonTalkTimeFilter `field:"optional" json:"nonTalkTimeFilter" yaml:"nonTalkTimeFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/transcribe_call_analytics_category#sentiment_filter TranscribeCallAnalyticsCategory#sentiment_filter}.
	SentimentFilter *TranscribeCallAnalyticsCategoryRulesSentimentFilter `field:"optional" json:"sentimentFilter" yaml:"sentimentFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/transcribe_call_analytics_category#transcript_filter TranscribeCallAnalyticsCategory#transcript_filter}.
	TranscriptFilter *TranscribeCallAnalyticsCategoryRulesTranscriptFilter `field:"optional" json:"transcriptFilter" yaml:"transcriptFilter"`
}

