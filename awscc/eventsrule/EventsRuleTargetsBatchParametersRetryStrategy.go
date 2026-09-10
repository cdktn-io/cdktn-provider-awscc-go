// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsrule


type EventsRuleTargetsBatchParametersRetryStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_rule#attempts EventsRule#attempts}.
	Attempts *float64 `field:"optional" json:"attempts" yaml:"attempts"`
}

