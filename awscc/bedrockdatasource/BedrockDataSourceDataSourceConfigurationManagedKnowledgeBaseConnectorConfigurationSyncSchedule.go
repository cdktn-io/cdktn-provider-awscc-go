// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncSchedule struct {
	// A daily refresh. The run time is system-chosen (off-peak) and not customer-configurable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/bedrock_data_source#daily BedrockDataSource#daily}
	Daily *string `field:"optional" json:"daily" yaml:"daily"`
	// A monthly refresh on a specified day of the month.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/bedrock_data_source#monthly BedrockDataSource#monthly}
	Monthly *BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleMonthly `field:"optional" json:"monthly" yaml:"monthly"`
	// A weekly refresh on a specified day of the week.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/bedrock_data_source#weekly BedrockDataSource#weekly}
	Weekly *BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleWeekly `field:"optional" json:"weekly" yaml:"weekly"`
}

