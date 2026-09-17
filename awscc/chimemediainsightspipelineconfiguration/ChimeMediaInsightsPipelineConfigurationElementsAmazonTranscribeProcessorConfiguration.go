// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimemediainsightspipelineconfiguration


type ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfiguration struct {
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
	// Turns language identification on or off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#identify_language ChimeMediaInsightsPipelineConfiguration#identify_language}
	IdentifyLanguage interface{} `field:"optional" json:"identifyLanguage" yaml:"identifyLanguage"`
	// Turns multiple language identification on or off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#identify_multiple_languages ChimeMediaInsightsPipelineConfiguration#identify_multiple_languages}
	IdentifyMultipleLanguages interface{} `field:"optional" json:"identifyMultipleLanguages" yaml:"identifyMultipleLanguages"`
	// The language code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#language_code ChimeMediaInsightsPipelineConfiguration#language_code}
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// The name of the custom language model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#language_model_name ChimeMediaInsightsPipelineConfiguration#language_model_name}
	LanguageModelName *string `field:"optional" json:"languageModelName" yaml:"languageModelName"`
	// The language options for transcription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#language_options ChimeMediaInsightsPipelineConfiguration#language_options}
	LanguageOptions *string `field:"optional" json:"languageOptions" yaml:"languageOptions"`
	// The level of stability for partial results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#partial_results_stability ChimeMediaInsightsPipelineConfiguration#partial_results_stability}
	PartialResultsStability *string `field:"optional" json:"partialResultsStability" yaml:"partialResultsStability"`
	// The types of PII to redact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#pii_entity_types ChimeMediaInsightsPipelineConfiguration#pii_entity_types}
	PiiEntityTypes *string `field:"optional" json:"piiEntityTypes" yaml:"piiEntityTypes"`
	// The preferred language for transcription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#preferred_language ChimeMediaInsightsPipelineConfiguration#preferred_language}
	PreferredLanguage *string `field:"optional" json:"preferredLanguage" yaml:"preferredLanguage"`
	// Enables speaker partitioning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#show_speaker_label ChimeMediaInsightsPipelineConfiguration#show_speaker_label}
	ShowSpeakerLabel interface{} `field:"optional" json:"showSpeakerLabel" yaml:"showSpeakerLabel"`
	// The vocabulary filtering method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#vocabulary_filter_method ChimeMediaInsightsPipelineConfiguration#vocabulary_filter_method}
	VocabularyFilterMethod *string `field:"optional" json:"vocabularyFilterMethod" yaml:"vocabularyFilterMethod"`
	// The name of the custom vocabulary filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#vocabulary_filter_name ChimeMediaInsightsPipelineConfiguration#vocabulary_filter_name}
	VocabularyFilterName *string `field:"optional" json:"vocabularyFilterName" yaml:"vocabularyFilterName"`
	// The names of the custom vocabulary filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#vocabulary_filter_names ChimeMediaInsightsPipelineConfiguration#vocabulary_filter_names}
	VocabularyFilterNames *string `field:"optional" json:"vocabularyFilterNames" yaml:"vocabularyFilterNames"`
	// The name of the custom vocabulary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#vocabulary_name ChimeMediaInsightsPipelineConfiguration#vocabulary_name}
	VocabularyName *string `field:"optional" json:"vocabularyName" yaml:"vocabularyName"`
	// The names of the custom vocabularies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#vocabulary_names ChimeMediaInsightsPipelineConfiguration#vocabulary_names}
	VocabularyNames *string `field:"optional" json:"vocabularyNames" yaml:"vocabularyNames"`
}

