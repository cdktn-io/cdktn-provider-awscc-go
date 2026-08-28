// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentagentspace

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DevopsagentAgentSpaceConfig struct {
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
	// The name of the AgentSpace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_agent_space#name DevopsagentAgentSpace#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the AgentSpace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_agent_space#description DevopsagentAgentSpace#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ARN of the KMS key to use for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_agent_space#kms_key_arn DevopsagentAgentSpace#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The locale for the AgentSpace, which determines the language used in agent responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_agent_space#locale DevopsagentAgentSpace#locale}
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_agent_space#operator_app DevopsagentAgentSpace#operator_app}.
	OperatorApp *DevopsagentAgentSpaceOperatorApp `field:"optional" json:"operatorApp" yaml:"operatorApp"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_agent_space#tags DevopsagentAgentSpace#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

