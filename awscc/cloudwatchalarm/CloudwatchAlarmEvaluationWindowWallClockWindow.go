// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalarm


type CloudwatchAlarmEvaluationWindowWallClockWindow struct {
	// The timezone for wall clock evaluation, in IANA time zone format (e.g., America/New_York, UTC).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudwatch_alarm#timezone CloudwatchAlarm#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

