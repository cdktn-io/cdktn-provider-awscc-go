// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transcribecallanalyticscategory


type TranscribeCallAnalyticsCategoryRulesTranscriptFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/transcribe_call_analytics_category#absolute_time_range TranscribeCallAnalyticsCategory#absolute_time_range}.
	AbsoluteTimeRange *TranscribeCallAnalyticsCategoryRulesTranscriptFilterAbsoluteTimeRange `field:"optional" json:"absoluteTimeRange" yaml:"absoluteTimeRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/transcribe_call_analytics_category#negate TranscribeCallAnalyticsCategory#negate}.
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/transcribe_call_analytics_category#participant_role TranscribeCallAnalyticsCategory#participant_role}.
	ParticipantRole *string `field:"optional" json:"participantRole" yaml:"participantRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/transcribe_call_analytics_category#relative_time_range TranscribeCallAnalyticsCategory#relative_time_range}.
	RelativeTimeRange *TranscribeCallAnalyticsCategoryRulesTranscriptFilterRelativeTimeRange `field:"optional" json:"relativeTimeRange" yaml:"relativeTimeRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/transcribe_call_analytics_category#targets TranscribeCallAnalyticsCategory#targets}.
	Targets *[]*string `field:"optional" json:"targets" yaml:"targets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/transcribe_call_analytics_category#transcript_filter_type TranscribeCallAnalyticsCategory#transcript_filter_type}.
	TranscriptFilterType *string `field:"optional" json:"transcriptFilterType" yaml:"transcriptFilterType"`
}

