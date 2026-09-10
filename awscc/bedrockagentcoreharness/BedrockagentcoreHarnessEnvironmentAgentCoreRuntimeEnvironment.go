// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessEnvironmentAgentCoreRuntimeEnvironment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_harness#filesystem_configurations BedrockagentcoreHarness#filesystem_configurations}.
	FilesystemConfigurations interface{} `field:"optional" json:"filesystemConfigurations" yaml:"filesystemConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_harness#lifecycle_configuration BedrockagentcoreHarness#lifecycle_configuration}.
	LifecycleConfiguration *BedrockagentcoreHarnessEnvironmentAgentCoreRuntimeEnvironmentLifecycleConfiguration `field:"optional" json:"lifecycleConfiguration" yaml:"lifecycleConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_harness#network_configuration BedrockagentcoreHarness#network_configuration}.
	NetworkConfiguration *BedrockagentcoreHarnessEnvironmentAgentCoreRuntimeEnvironmentNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
}

