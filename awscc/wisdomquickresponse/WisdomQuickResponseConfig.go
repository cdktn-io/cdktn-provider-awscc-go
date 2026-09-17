// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomquickresponse

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WisdomQuickResponseConfig struct {
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
	// The container of quick response content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_quick_response#content WisdomQuickResponse#content}
	Content *WisdomQuickResponseContent `field:"required" json:"content" yaml:"content"`
	// The Amazon Resource Name (ARN) of the knowledge base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_quick_response#knowledge_base_arn WisdomQuickResponse#knowledge_base_arn}
	KnowledgeBaseArn *string `field:"required" json:"knowledgeBaseArn" yaml:"knowledgeBaseArn"`
	// The name of the quick response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_quick_response#name WisdomQuickResponse#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Connect contact channels this quick response applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_quick_response#channels WisdomQuickResponse#channels}
	Channels *[]*string `field:"optional" json:"channels" yaml:"channels"`
	// The media type of the quick response content.
	//
	// - Use application/x.quickresponse;format=plain for quick response written in plain text.
	// - Use application/x.quickresponse;format=markdown for quick response written in richtext.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_quick_response#content_type WisdomQuickResponse#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The description of the quick response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_quick_response#description WisdomQuickResponse#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The configuration information of the user groups that the quick response is accessible to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_quick_response#grouping_configuration WisdomQuickResponse#grouping_configuration}
	GroupingConfiguration *WisdomQuickResponseGroupingConfiguration `field:"optional" json:"groupingConfiguration" yaml:"groupingConfiguration"`
	// Whether the quick response is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_quick_response#is_active WisdomQuickResponse#is_active}
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
	// The language code value for the language in which the quick response is written.
	//
	// The supported language codes include de_DE, en_US, es_ES, fr_FR, id_ID, it_IT, ja_JP, ko_KR, pt_BR, zh_CN, zh_TW
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_quick_response#language WisdomQuickResponse#language}
	Language *string `field:"optional" json:"language" yaml:"language"`
	// The shortcut key of the quick response. The value should be unique across the knowledge base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_quick_response#shortcut_key WisdomQuickResponse#shortcut_key}
	ShortcutKey *string `field:"optional" json:"shortcutKey" yaml:"shortcutKey"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_quick_response#tags WisdomQuickResponse#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

