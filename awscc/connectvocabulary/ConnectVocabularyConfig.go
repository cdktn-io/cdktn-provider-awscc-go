// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectvocabulary

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectVocabularyConfig struct {
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
	// The content of the custom vocabulary in plain-text format with a table of values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_vocabulary#content ConnectVocabulary#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The identifier of the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_vocabulary#instance_id ConnectVocabulary#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The language code of the vocabulary entries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_vocabulary#language_code ConnectVocabulary#language_code}
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// A unique name of the custom vocabulary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_vocabulary#vocabulary_name ConnectVocabulary#vocabulary_name}
	VocabularyName *string `field:"required" json:"vocabularyName" yaml:"vocabularyName"`
	// The tags used to organize, track, or control access for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_vocabulary#tags ConnectVocabulary#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

