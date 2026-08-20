// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessEnvironmentArtifact struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#container_configuration BedrockagentcoreHarness#container_configuration}.
	ContainerConfiguration *BedrockagentcoreHarnessEnvironmentArtifactContainerConfiguration `field:"optional" json:"containerConfiguration" yaml:"containerConfiguration"`
}

