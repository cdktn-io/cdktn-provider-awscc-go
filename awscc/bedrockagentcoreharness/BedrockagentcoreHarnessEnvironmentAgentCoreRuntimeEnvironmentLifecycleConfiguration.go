// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessEnvironmentAgentCoreRuntimeEnvironmentLifecycleConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#idle_runtime_session_timeout BedrockagentcoreHarness#idle_runtime_session_timeout}.
	IdleRuntimeSessionTimeout *float64 `field:"optional" json:"idleRuntimeSessionTimeout" yaml:"idleRuntimeSessionTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#max_lifetime BedrockagentcoreHarness#max_lifetime}.
	MaxLifetime *float64 `field:"optional" json:"maxLifetime" yaml:"maxLifetime"`
}

