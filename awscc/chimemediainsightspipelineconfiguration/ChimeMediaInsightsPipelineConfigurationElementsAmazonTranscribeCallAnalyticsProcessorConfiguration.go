// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimemediainsightspipelineconfiguration


type ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfiguration struct {
	// The categories to send to the insights target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#call_analytics_stream_categories ChimeMediaInsightsPipelineConfiguration#call_analytics_stream_categories}
	CallAnalyticsStreamCategories *[]*string `field:"optional" json:"callAnalyticsStreamCategories" yaml:"callAnalyticsStreamCategories"`
	// Labels all PII identified in the transcript.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#content_identification_type ChimeMediaInsightsPipelineConfiguration#content_identification_type}
	ContentIdentificationType *string `field:"optional" json:"contentIdentificationType" yaml:"contentIdentificationType"`
	// Redacts all PII identified in the transcript.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#content_redaction_type ChimeMediaInsightsPipelineConfiguration#content_redaction_type}
	ContentRedactionType *string `field:"optional" json:"contentRedactionType" yaml:"contentRedactionType"`
	// Enables partial result stabilization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#enable_partial_results_stabilization ChimeMediaInsightsPipelineConfiguration#enable_partial_results_stabilization}
	EnablePartialResultsStabilization interface{} `field:"optional" json:"enablePartialResultsStabilization" yaml:"enablePartialResultsStabilization"`
	// If true, partial results are filtered out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#filter_partial_results ChimeMediaInsightsPipelineConfiguration#filter_partial_results}
	FilterPartialResults interface{} `field:"optional" json:"filterPartialResults" yaml:"filterPartialResults"`
	// The language code in the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#language_code ChimeMediaInsightsPipelineConfiguration#language_code}
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// The name of the custom language model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#language_model_name ChimeMediaInsightsPipelineConfiguration#language_model_name}
	LanguageModelName *string `field:"optional" json:"languageModelName" yaml:"languageModelName"`
	// The level of stability for partial results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#partial_results_stability ChimeMediaInsightsPipelineConfiguration#partial_results_stability}
	PartialResultsStability *string `field:"optional" json:"partialResultsStability" yaml:"partialResultsStability"`
	// The types of PII to redact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#pii_entity_types ChimeMediaInsightsPipelineConfiguration#pii_entity_types}
	PiiEntityTypes *string `field:"optional" json:"piiEntityTypes" yaml:"piiEntityTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#post_call_analytics_settings ChimeMediaInsightsPipelineConfiguration#post_call_analytics_settings}.
	PostCallAnalyticsSettings *ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationPostCallAnalyticsSettings `field:"optional" json:"postCallAnalyticsSettings" yaml:"postCallAnalyticsSettings"`
	// The vocabulary filtering method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#vocabulary_filter_method ChimeMediaInsightsPipelineConfiguration#vocabulary_filter_method}
	VocabularyFilterMethod *string `field:"optional" json:"vocabularyFilterMethod" yaml:"vocabularyFilterMethod"`
	// The name of the custom vocabulary filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#vocabulary_filter_name ChimeMediaInsightsPipelineConfiguration#vocabulary_filter_name}
	VocabularyFilterName *string `field:"optional" json:"vocabularyFilterName" yaml:"vocabularyFilterName"`
	// The name of the custom vocabulary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#vocabulary_name ChimeMediaInsightsPipelineConfiguration#vocabulary_name}
	VocabularyName *string `field:"optional" json:"vocabularyName" yaml:"vocabularyName"`
}

