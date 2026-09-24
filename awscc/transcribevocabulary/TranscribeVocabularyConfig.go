// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transcribevocabulary

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TranscribeVocabularyConfig struct {
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
	// The language code that represents the language of the entries in your custom vocabulary.
	//
	// Each custom vocabulary must contain terms in only one language.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/transcribe_vocabulary#language_code TranscribeVocabulary#language_code}
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// A unique name, chosen by you, for your custom vocabulary.
	//
	// This name is case sensitive, cannot contain spaces, and must be unique within an AWS account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/transcribe_vocabulary#vocabulary_name TranscribeVocabulary#vocabulary_name}
	VocabularyName *string `field:"required" json:"vocabularyName" yaml:"vocabularyName"`
	// The Amazon Resource Name (ARN) of an IAM role that has permissions to access the Amazon S3 bucket that contains your input file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/transcribe_vocabulary#data_access_role_arn TranscribeVocabulary#data_access_role_arn}
	DataAccessRoleArn *string `field:"optional" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// Use this parameter if you want to create your custom vocabulary by including all desired terms, as comma-separated values, within your request.
	//
	// You cannot use this parameter together with VocabularyFileUri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/transcribe_vocabulary#phrases TranscribeVocabulary#phrases}
	Phrases *[]*string `field:"optional" json:"phrases" yaml:"phrases"`
	// Adds one or more custom tags, each in the form of a key:value pair, to the custom vocabulary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/transcribe_vocabulary#tags TranscribeVocabulary#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon S3 location of the text file that contains your custom vocabulary.
	//
	// You cannot use this parameter together with Phrases.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/transcribe_vocabulary#vocabulary_file_uri TranscribeVocabulary#vocabulary_file_uri}
	VocabularyFileUri *string `field:"optional" json:"vocabularyFileUri" yaml:"vocabularyFileUri"`
}

