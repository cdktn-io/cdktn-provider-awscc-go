// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdommessagetemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WisdomMessageTemplateConfig struct {
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
	// The channel subtype this message template applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_message_template#channel_subtype WisdomMessageTemplate#channel_subtype}
	ChannelSubtype *string `field:"required" json:"channelSubtype" yaml:"channelSubtype"`
	// The content of the message template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_message_template#content WisdomMessageTemplate#content}
	Content *WisdomMessageTemplateContent `field:"required" json:"content" yaml:"content"`
	// The Amazon Resource Name (ARN) of the knowledge base to which the message template belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_message_template#knowledge_base_arn WisdomMessageTemplate#knowledge_base_arn}
	KnowledgeBaseArn *string `field:"required" json:"knowledgeBaseArn" yaml:"knowledgeBaseArn"`
	// The name of the message template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_message_template#name WisdomMessageTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An object that specifies the default values to use for variables in the message template.
	//
	// This object contains different categories of key-value pairs. Each key defines a variable or placeholder in the message template. The corresponding value defines the default value for that variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_message_template#default_attributes WisdomMessageTemplate#default_attributes}
	DefaultAttributes *WisdomMessageTemplateDefaultAttributes `field:"optional" json:"defaultAttributes" yaml:"defaultAttributes"`
	// The description of the message template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_message_template#description WisdomMessageTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The configuration information of the user groups that the message template is accessible to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_message_template#grouping_configuration WisdomMessageTemplate#grouping_configuration}
	GroupingConfiguration *WisdomMessageTemplateGroupingConfiguration `field:"optional" json:"groupingConfiguration" yaml:"groupingConfiguration"`
	// The language code value for the language in which the message template is written.
	//
	// The supported language codes include de_DE, en_US, es_ES, fr_FR, id_ID, it_IT, ja_JP, ko_KR, pt_BR, zh_CN, zh_TW
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_message_template#language WisdomMessageTemplate#language}
	Language *string `field:"optional" json:"language" yaml:"language"`
	// List of message template attachments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_message_template#message_template_attachments WisdomMessageTemplate#message_template_attachments}
	MessageTemplateAttachments interface{} `field:"optional" json:"messageTemplateAttachments" yaml:"messageTemplateAttachments"`
	// The tags used to organize, track, or control access for this resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_message_template#tags WisdomMessageTemplate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

