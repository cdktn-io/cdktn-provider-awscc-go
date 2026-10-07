// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transcribemedicalvocabulary

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TranscribeMedicalVocabularyConfig struct {
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
	// The language code of the vocabulary entries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/transcribe_medical_vocabulary#language_code TranscribeMedicalVocabulary#language_code}
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// The name of the medical vocabulary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/transcribe_medical_vocabulary#vocabulary_name TranscribeMedicalVocabulary#vocabulary_name}
	VocabularyName *string `field:"required" json:"vocabularyName" yaml:"vocabularyName"`
	// Tags associated with the medical vocabulary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/transcribe_medical_vocabulary#tags TranscribeMedicalVocabulary#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon S3 location of the text file that contains the medical vocabulary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/transcribe_medical_vocabulary#vocabulary_file_uri TranscribeMedicalVocabulary#vocabulary_file_uri}
	VocabularyFileUri *string `field:"optional" json:"vocabularyFileUri" yaml:"vocabularyFileUri"`
}

