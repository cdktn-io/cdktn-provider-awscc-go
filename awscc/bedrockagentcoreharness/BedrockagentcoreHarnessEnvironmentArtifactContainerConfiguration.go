// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessEnvironmentArtifactContainerConfiguration struct {
	// The ECR URI of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_harness#container_uri BedrockagentcoreHarness#container_uri}
	ContainerUri *string `field:"optional" json:"containerUri" yaml:"containerUri"`
}

