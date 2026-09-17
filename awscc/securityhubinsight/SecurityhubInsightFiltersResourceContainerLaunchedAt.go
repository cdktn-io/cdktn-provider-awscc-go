// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubinsight


type SecurityhubInsightFiltersResourceContainerLaunchedAt struct {
	// A date range for the date filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/securityhub_insight#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersResourceContainerLaunchedAtDateRange `field:"optional" json:"dateRange" yaml:"dateRange"`
	// The date and time, in UTC and ISO 8601 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/securityhub_insight#end SecurityhubInsight#end}
	End *string `field:"optional" json:"end" yaml:"end"`
	// The date and time, in UTC and ISO 8601 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/securityhub_insight#start SecurityhubInsight#start}
	Start *string `field:"optional" json:"start" yaml:"start"`
}

