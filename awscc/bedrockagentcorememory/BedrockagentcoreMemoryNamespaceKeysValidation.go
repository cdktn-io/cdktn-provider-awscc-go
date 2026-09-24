// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememory


type BedrockagentcoreMemoryNamespaceKeysValidation struct {
	// List of allowed values for this namespace variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_memory#allowed_values BedrockagentcoreMemory#allowed_values}
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// A regex pattern that a namespace variable value must match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_memory#regex_pattern BedrockagentcoreMemory#regex_pattern}
	RegexPattern *string `field:"optional" json:"regexPattern" yaml:"regexPattern"`
}

