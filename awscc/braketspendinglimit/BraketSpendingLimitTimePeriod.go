// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package braketspendinglimit


type BraketSpendingLimitTimePeriod struct {
	// The end date and time for the spending limit period, in ISO 8601 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/braket_spending_limit#end_at BraketSpendingLimit#end_at}
	EndAt *string `field:"optional" json:"endAt" yaml:"endAt"`
	// The start date and time for the spending limit period, in ISO 8601 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/braket_spending_limit#start_at BraketSpendingLimit#start_at}
	StartAt *string `field:"optional" json:"startAt" yaml:"startAt"`
}

