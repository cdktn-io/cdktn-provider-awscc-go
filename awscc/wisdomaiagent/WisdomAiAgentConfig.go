// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WisdomAiAgentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_ai_agent#assistant_id WisdomAiAgent#assistant_id}.
	AssistantId *string `field:"required" json:"assistantId" yaml:"assistantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_ai_agent#configuration WisdomAiAgent#configuration}.
	Configuration *WisdomAiAgentConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_ai_agent#type WisdomAiAgent#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_ai_agent#description WisdomAiAgent#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_ai_agent#name WisdomAiAgent#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_ai_agent#tags WisdomAiAgent#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

