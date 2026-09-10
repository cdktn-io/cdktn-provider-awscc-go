// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotTestBotAliasSettingsConversationLogSettingsAudioLogSettingsDestination struct {
	// Specifies an Amazon S3 bucket for logging audio conversations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lex_bot#s3_bucket LexBot#s3_bucket}
	S3Bucket *LexBotTestBotAliasSettingsConversationLogSettingsAudioLogSettingsDestinationS3Bucket `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
}

