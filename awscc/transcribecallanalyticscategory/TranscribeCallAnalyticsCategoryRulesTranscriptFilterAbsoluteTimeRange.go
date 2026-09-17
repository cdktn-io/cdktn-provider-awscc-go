// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transcribecallanalyticscategory


type TranscribeCallAnalyticsCategoryRulesTranscriptFilterAbsoluteTimeRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/transcribe_call_analytics_category#end_time TranscribeCallAnalyticsCategory#end_time}.
	EndTime *float64 `field:"optional" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/transcribe_call_analytics_category#first TranscribeCallAnalyticsCategory#first}.
	First *float64 `field:"optional" json:"first" yaml:"first"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/transcribe_call_analytics_category#last TranscribeCallAnalyticsCategory#last}.
	Last *float64 `field:"optional" json:"last" yaml:"last"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/transcribe_call_analytics_category#start_time TranscribeCallAnalyticsCategory#start_time}.
	StartTime *float64 `field:"optional" json:"startTime" yaml:"startTime"`
}

