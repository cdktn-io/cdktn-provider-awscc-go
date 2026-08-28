// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arczonalshiftzonalautoshiftconfiguration


type ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/arczonalshift_zonal_autoshift_configuration#blocked_dates ArczonalshiftZonalAutoshiftConfiguration#blocked_dates}.
	BlockedDates *[]*string `field:"optional" json:"blockedDates" yaml:"blockedDates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/arczonalshift_zonal_autoshift_configuration#blocked_windows ArczonalshiftZonalAutoshiftConfiguration#blocked_windows}.
	BlockedWindows *[]*string `field:"optional" json:"blockedWindows" yaml:"blockedWindows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/arczonalshift_zonal_autoshift_configuration#blocking_alarms ArczonalshiftZonalAutoshiftConfiguration#blocking_alarms}.
	BlockingAlarms interface{} `field:"optional" json:"blockingAlarms" yaml:"blockingAlarms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/arczonalshift_zonal_autoshift_configuration#outcome_alarms ArczonalshiftZonalAutoshiftConfiguration#outcome_alarms}.
	OutcomeAlarms interface{} `field:"optional" json:"outcomeAlarms" yaml:"outcomeAlarms"`
}

