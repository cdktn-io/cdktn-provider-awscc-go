// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagent


type BedrockAgentMemoryConfiguration struct {
	// Types of session storage persisted in memory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_agent#enabled_memory_types BedrockAgent#enabled_memory_types}
	EnabledMemoryTypes *[]*string `field:"optional" json:"enabledMemoryTypes" yaml:"enabledMemoryTypes"`
	// Configuration for Session Summarization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_agent#session_summary_configuration BedrockAgent#session_summary_configuration}
	SessionSummaryConfiguration *BedrockAgentMemoryConfigurationSessionSummaryConfiguration `field:"optional" json:"sessionSummaryConfiguration" yaml:"sessionSummaryConfiguration"`
	// Maximum number of days to store session details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_agent#storage_days BedrockAgent#storage_days}
	StorageDays *float64 `field:"optional" json:"storageDays" yaml:"storageDays"`
}

