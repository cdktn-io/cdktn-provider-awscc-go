// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorecodeinterpretercustom

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreCodeInterpreterCustomConfig struct {
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
	// The name of the code interpreter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_code_interpreter_custom#name BedrockagentcoreCodeInterpreterCustom#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Network configuration for code interpreter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_code_interpreter_custom#network_configuration BedrockagentcoreCodeInterpreterCustom#network_configuration}
	NetworkConfiguration *BedrockagentcoreCodeInterpreterCustomNetworkConfiguration `field:"required" json:"networkConfiguration" yaml:"networkConfiguration"`
	// List of root CA certificates in PEM format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_code_interpreter_custom#certificates BedrockagentcoreCodeInterpreterCustom#certificates}
	Certificates interface{} `field:"optional" json:"certificates" yaml:"certificates"`
	// The description of the code interpreter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_code_interpreter_custom#description BedrockagentcoreCodeInterpreterCustom#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ARN of the IAM role that the code interpreter uses to access resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_code_interpreter_custom#execution_role_arn BedrockagentcoreCodeInterpreterCustom#execution_role_arn}
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// A map of tag keys and values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_code_interpreter_custom#tags BedrockagentcoreCodeInterpreterCustom#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

