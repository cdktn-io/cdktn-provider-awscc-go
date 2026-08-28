// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreruntime


type BedrockagentcoreRuntimeRequestHeaderConfiguration struct {
	// List of allowed HTTP headers for agent runtime requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_runtime#request_header_allowlist BedrockagentcoreRuntime#request_header_allowlist}
	RequestHeaderAllowlist *[]*string `field:"optional" json:"requestHeaderAllowlist" yaml:"requestHeaderAllowlist"`
}

