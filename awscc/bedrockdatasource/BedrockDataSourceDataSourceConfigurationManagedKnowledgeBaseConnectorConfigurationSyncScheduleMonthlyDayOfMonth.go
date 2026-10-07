// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationSyncScheduleMonthlyDayOfMonth struct {
	// Specific day of the month, 1 through 28 (capped at 28 to avoid month-length ambiguity).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/bedrock_data_source#day_number BedrockDataSource#day_number}
	DayNumber *float64 `field:"optional" json:"dayNumber" yaml:"dayNumber"`
	// Run on the last calendar day of each month.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/bedrock_data_source#last_day_of_month BedrockDataSource#last_day_of_month}
	LastDayOfMonth *string `field:"optional" json:"lastDayOfMonth" yaml:"lastDayOfMonth"`
}

