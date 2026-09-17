// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimemediainsightspipelineconfiguration


type ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationPostCallAnalyticsSettings struct {
	// The content redaction output settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#content_redaction_output ChimeMediaInsightsPipelineConfiguration#content_redaction_output}
	ContentRedactionOutput *string `field:"optional" json:"contentRedactionOutput" yaml:"contentRedactionOutput"`
	// The ARN of the role used by Transcribe to upload post-call analysis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#data_access_role_arn ChimeMediaInsightsPipelineConfiguration#data_access_role_arn}
	DataAccessRoleArn *string `field:"optional" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// The ID of the KMS key used to encrypt the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#output_encryption_kms_key_id ChimeMediaInsightsPipelineConfiguration#output_encryption_kms_key_id}
	OutputEncryptionKmsKeyId *string `field:"optional" json:"outputEncryptionKmsKeyId" yaml:"outputEncryptionKmsKeyId"`
	// The URL of the Amazon S3 bucket for post-call data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#output_location ChimeMediaInsightsPipelineConfiguration#output_location}
	OutputLocation *string `field:"optional" json:"outputLocation" yaml:"outputLocation"`
}

