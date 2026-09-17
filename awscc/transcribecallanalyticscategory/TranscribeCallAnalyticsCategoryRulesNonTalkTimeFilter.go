// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transcribecallanalyticscategory


type TranscribeCallAnalyticsCategoryRulesNonTalkTimeFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/transcribe_call_analytics_category#absolute_time_range TranscribeCallAnalyticsCategory#absolute_time_range}.
	AbsoluteTimeRange *TranscribeCallAnalyticsCategoryRulesNonTalkTimeFilterAbsoluteTimeRange `field:"optional" json:"absoluteTimeRange" yaml:"absoluteTimeRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/transcribe_call_analytics_category#negate TranscribeCallAnalyticsCategory#negate}.
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/transcribe_call_analytics_category#relative_time_range TranscribeCallAnalyticsCategory#relative_time_range}.
	RelativeTimeRange *TranscribeCallAnalyticsCategoryRulesNonTalkTimeFilterRelativeTimeRange `field:"optional" json:"relativeTimeRange" yaml:"relativeTimeRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/transcribe_call_analytics_category#threshold TranscribeCallAnalyticsCategory#threshold}.
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
}

