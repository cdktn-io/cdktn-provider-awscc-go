// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecustomentitytype

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueCustomEntityTypeConfig struct {
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
	// A list of context words.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_custom_entity_type#context_words GlueCustomEntityType#context_words}
	ContextWords *[]*string `field:"optional" json:"contextWords" yaml:"contextWords"`
	// The name of the custom entity type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_custom_entity_type#name GlueCustomEntityType#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A regular expression string that is used for detecting sensitive data in a custom pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_custom_entity_type#regex_string GlueCustomEntityType#regex_string}
	RegexString *string `field:"optional" json:"regexString" yaml:"regexString"`
	// Tags to associate with the custom entity type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_custom_entity_type#tags GlueCustomEntityType#tags}
	Tags *string `field:"optional" json:"tags" yaml:"tags"`
}

