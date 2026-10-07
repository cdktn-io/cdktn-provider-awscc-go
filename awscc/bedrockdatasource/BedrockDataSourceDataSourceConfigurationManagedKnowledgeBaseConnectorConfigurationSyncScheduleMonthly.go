// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleMonthly struct {
	// Day of the month on which a monthly refresh runs.
	//
	// Exactly one variant is set: an explicit day number, or the last calendar day of the month.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/bedrock_data_source#day_of_month BedrockDataSource#day_of_month}
	DayOfMonth *BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleMonthlyDayOfMonth `field:"optional" json:"dayOfMonth" yaml:"dayOfMonth"`
}

