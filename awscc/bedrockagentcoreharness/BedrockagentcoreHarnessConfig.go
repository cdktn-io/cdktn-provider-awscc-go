// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreHarnessConfig struct {
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
	// The ARN of the IAM role that the harness assumes when running.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#execution_role_arn BedrockagentcoreHarness#execution_role_arn}
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The name of the harness.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#harness_name BedrockagentcoreHarness#harness_name}
	HarnessName *string `field:"required" json:"harnessName" yaml:"harnessName"`
	// The model configuration for the harness.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#model BedrockagentcoreHarness#model}
	Model *BedrockagentcoreHarnessModel `field:"required" json:"model" yaml:"model"`
	// The tools that the agent is allowed to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#allowed_tools BedrockagentcoreHarness#allowed_tools}
	AllowedTools *[]*string `field:"optional" json:"allowedTools" yaml:"allowedTools"`
	// The inbound authorization configuration for authenticating incoming requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#authorizer_configuration BedrockagentcoreHarness#authorizer_configuration}
	AuthorizerConfiguration *BedrockagentcoreHarnessAuthorizerConfiguration `field:"optional" json:"authorizerConfiguration" yaml:"authorizerConfiguration"`
	// The compute environment configuration for the harness, including underlying runtime information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#environment BedrockagentcoreHarness#environment}
	Environment *BedrockagentcoreHarnessEnvironment `field:"optional" json:"environment" yaml:"environment"`
	// The environment artifact for the harness, such as a custom container image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#environment_artifact BedrockagentcoreHarness#environment_artifact}
	EnvironmentArtifact *BedrockagentcoreHarnessEnvironmentArtifact `field:"optional" json:"environmentArtifact" yaml:"environmentArtifact"`
	// Environment variables to set in the harness runtime environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#environment_variables BedrockagentcoreHarness#environment_variables}
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// The maximum number of iterations the agent loop can execute per invocation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#max_iterations BedrockagentcoreHarness#max_iterations}
	MaxIterations *float64 `field:"optional" json:"maxIterations" yaml:"maxIterations"`
	// The maximum number of tokens the agent can generate per iteration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#max_tokens BedrockagentcoreHarness#max_tokens}
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// The AgentCore Memory configuration for persisting conversation context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#memory BedrockagentcoreHarness#memory}
	Memory *BedrockagentcoreHarnessMemory `field:"optional" json:"memory" yaml:"memory"`
	// The skills available to the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#skills BedrockagentcoreHarness#skills}
	Skills interface{} `field:"optional" json:"skills" yaml:"skills"`
	// The system prompt that defines the agent's behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#system_prompt BedrockagentcoreHarness#system_prompt}
	SystemPrompt interface{} `field:"optional" json:"systemPrompt" yaml:"systemPrompt"`
	// Tags to apply to the harness resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#tags BedrockagentcoreHarness#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The maximum duration in seconds for the agent loop execution per invocation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#timeout_seconds BedrockagentcoreHarness#timeout_seconds}
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// The tools available to the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#tools BedrockagentcoreHarness#tools}
	Tools interface{} `field:"optional" json:"tools" yaml:"tools"`
	// The truncation configuration for managing conversation context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#truncation BedrockagentcoreHarness#truncation}
	Truncation *BedrockagentcoreHarnessTruncation `field:"optional" json:"truncation" yaml:"truncation"`
}

