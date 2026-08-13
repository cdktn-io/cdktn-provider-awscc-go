// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transcribevocabularyfilter

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TranscribeVocabularyFilterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The language code that represents the language of the entries in your vocabulary filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/transcribe_vocabulary_filter#language_code TranscribeVocabularyFilter#language_code}
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// A unique name, chosen by you, for your custom vocabulary filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/transcribe_vocabulary_filter#vocabulary_filter_name TranscribeVocabularyFilter#vocabulary_filter_name}
	VocabularyFilterName *string `field:"required" json:"vocabularyFilterName" yaml:"vocabularyFilterName"`
	// The Amazon Resource Name (ARN) of an IAM role that has permissions to access the Amazon S3 bucket that contains your input files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/transcribe_vocabulary_filter#data_access_role_arn TranscribeVocabularyFilter#data_access_role_arn}
	DataAccessRoleArn *string `field:"optional" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// Tags associated with the vocabulary filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/transcribe_vocabulary_filter#tags TranscribeVocabularyFilter#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon S3 location of the text file that contains your custom vocabulary filter terms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/transcribe_vocabulary_filter#vocabulary_filter_file_uri TranscribeVocabularyFilter#vocabulary_filter_file_uri}
	VocabularyFilterFileUri *string `field:"optional" json:"vocabularyFilterFileUri" yaml:"vocabularyFilterFileUri"`
	// Use this parameter if you want to create your custom vocabulary filter by including all desired terms, as comma-separated values, within your request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/transcribe_vocabulary_filter#words TranscribeVocabularyFilter#words}
	Words *[]*string `field:"optional" json:"words" yaml:"words"`
}

