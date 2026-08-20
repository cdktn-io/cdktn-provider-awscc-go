// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememory


type BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_memory#message_based_trigger BedrockagentcoreMemory#message_based_trigger}.
	MessageBasedTrigger *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTrigger `field:"optional" json:"messageBasedTrigger" yaml:"messageBasedTrigger"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_memory#time_based_trigger BedrockagentcoreMemory#time_based_trigger}.
	TimeBasedTrigger *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsTimeBasedTrigger `field:"optional" json:"timeBasedTrigger" yaml:"timeBasedTrigger"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_memory#token_based_trigger BedrockagentcoreMemory#token_based_trigger}.
	TokenBasedTrigger *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsTokenBasedTrigger `field:"optional" json:"tokenBasedTrigger" yaml:"tokenBasedTrigger"`
}

