// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreruntime


type BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerAllowedWorkloadConfiguration struct {
	// List of allow-listed hosting environments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_runtime#hosting_environments BedrockagentcoreRuntime#hosting_environments}
	HostingEnvironments interface{} `field:"optional" json:"hostingEnvironments" yaml:"hostingEnvironments"`
	// List of allow-listed workload identity names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_runtime#workload_identities BedrockagentcoreRuntime#workload_identities}
	WorkloadIdentities *[]*string `field:"optional" json:"workloadIdentities" yaml:"workloadIdentities"`
}

