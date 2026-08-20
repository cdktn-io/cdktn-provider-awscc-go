// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagent


type BedrockAgentMemoryConfigurationSessionSummaryConfiguration struct {
	// Maximum number of Sessions to Summarize.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_agent#max_recent_sessions BedrockAgent#max_recent_sessions}
	MaxRecentSessions *float64 `field:"optional" json:"maxRecentSessions" yaml:"maxRecentSessions"`
}

