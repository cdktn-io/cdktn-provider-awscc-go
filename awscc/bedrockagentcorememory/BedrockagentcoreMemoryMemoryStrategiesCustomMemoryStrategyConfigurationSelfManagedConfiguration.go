// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememory


type BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSelfManagedConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_memory#historical_context_window_size BedrockagentcoreMemory#historical_context_window_size}.
	HistoricalContextWindowSize *float64 `field:"optional" json:"historicalContextWindowSize" yaml:"historicalContextWindowSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_memory#invocation_configuration BedrockagentcoreMemory#invocation_configuration}.
	InvocationConfiguration *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSelfManagedConfigurationInvocationConfiguration `field:"optional" json:"invocationConfiguration" yaml:"invocationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_memory#trigger_conditions BedrockagentcoreMemory#trigger_conditions}.
	TriggerConditions interface{} `field:"optional" json:"triggerConditions" yaml:"triggerConditions"`
}

