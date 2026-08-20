// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdommessagetemplateversion

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WisdomMessageTemplateVersionConfig struct {
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
	// The unqualified Amazon Resource Name (ARN) of the message template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_message_template_version#message_template_arn WisdomMessageTemplateVersion#message_template_arn}
	MessageTemplateArn *string `field:"required" json:"messageTemplateArn" yaml:"messageTemplateArn"`
	// The content SHA256 of the message template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_message_template_version#message_template_content_sha_256 WisdomMessageTemplateVersion#message_template_content_sha_256}
	MessageTemplateContentSha256 *string `field:"optional" json:"messageTemplateContentSha256" yaml:"messageTemplateContentSha256"`
}

