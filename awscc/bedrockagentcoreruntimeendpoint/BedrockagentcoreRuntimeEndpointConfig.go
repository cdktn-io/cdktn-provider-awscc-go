// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreruntimeendpoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreRuntimeEndpointConfig struct {
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
	// The ID of the parent Agent Runtime (required for creation).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_runtime_endpoint#agent_runtime_id BedrockagentcoreRuntimeEndpoint#agent_runtime_id}
	AgentRuntimeId *string `field:"required" json:"agentRuntimeId" yaml:"agentRuntimeId"`
	// The name of the Agent Runtime Endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_runtime_endpoint#name BedrockagentcoreRuntimeEndpoint#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The version of the AgentCore Runtime to use for the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_runtime_endpoint#agent_runtime_version BedrockagentcoreRuntimeEndpoint#agent_runtime_version}
	AgentRuntimeVersion *string `field:"optional" json:"agentRuntimeVersion" yaml:"agentRuntimeVersion"`
	// The description of the AgentCore Runtime endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_runtime_endpoint#description BedrockagentcoreRuntimeEndpoint#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A map of tag keys and values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_runtime_endpoint#tags BedrockagentcoreRuntimeEndpoint#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

